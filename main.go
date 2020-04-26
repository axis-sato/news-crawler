package main

import (
	"os"
	"time"

	"github.com/c8112002/news-crawler/crawler"
	log "github.com/sirupsen/logrus"

	"github.com/c8112002/news-crawler/entities"

	"github.com/c8112002/news-crawler/db"
	"github.com/c8112002/news-crawler/store"
	"github.com/c8112002/news-crawler/utils"
	_ "github.com/go-sql-driver/mysql"
)

var now time.Time

func init() {
	now = time.Now()
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}

func main() {
	loadEnv()

	d, err := db.New(utils.GetEnv())
	if err != nil {
		panic(err.Error())
	}
	defer d.Close()

	ts := store.NewTagStore(d)
	ss := store.NewSiteStore(d)
	as := store.NewArticleStore(d)

	// クロール対象のタグを取得
	tags := getTags(ts)
	log.WithFields(log.Fields{
		"tags": tags,
	}).Debug("クロール対象のタグ")

	// Qiitaクロール
	qiitaResults := crawlQiita(tags)

	// Qiitaのアイテムを保存
	qiita := getQiita(ss)
	saveQiitaItem(qiitaResults, qiita, as)
}

func loadEnv() {
	if err := utils.LoadEnv(); err != nil {
		panic(err.Error())
	}
}

func getTags(ts *store.TagStore) *entities.Tags {
	tags, err := ts.GetFollowingTag()
	if err != nil {
		panic(err.Error())
	}

	return tags
}

func crawlQiita(tags *entities.Tags) []crawler.QiitaResult {
	today := now
	_1weekAgo := today.AddDate(0, 0, -7)
	ac := crawler.NewQiitaCrawler(
		os.Getenv("QIITA_TOKEN"),
		tags,
		_1weekAgo,
		today,
	)
	qiitaResults, err := ac.Run()
	if err != nil {
		panic(err.Error())
	}
	return qiitaResults
}

func getQiita(ss *store.SiteStore) *entities.Site {
	qiita, err := ss.GetQiita()
	if err != nil {
		panic(err.Error())
	}

	return qiita
}

func saveQiitaItem(
	qiitaResults []crawler.QiitaResult,
	qiita *entities.Site,
	as *store.ArticleStore,
) {
	for _, r := range qiitaResults {
		var articles []*entities.Article
		for _, item := range r.Items {
			a := entities.NewArticle(item.ID, item.Title, item.URL, item.Thumbnail, now)
			articles = append(articles, a)
		}
		res, err := as.SaveArticles(articles, r.Tag, qiita)
		if err != nil {
			panic(err.Error())
		}
		log.WithFields(log.Fields{
			"articles": res,
		}).Debug("保存したQiitaのアイテム")
	}
}
