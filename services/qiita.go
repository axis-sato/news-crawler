package services

import (
	"os"
	"time"

	"github.com/c8112002/news-crawler/crawler"
	"github.com/c8112002/news-crawler/entities"
	"github.com/c8112002/news-crawler/store"
	log "github.com/sirupsen/logrus"
)

type QiitaService struct {
	now  time.Time
	tags *entities.Tags
	ss   *store.SiteStore
	as   *store.ArticleStore
}

func NewQiitaService(now time.Time, tags *entities.Tags, ss *store.SiteStore, as *store.ArticleStore) *QiitaService {
	return &QiitaService{now: now, tags: tags, ss: ss, as: as}
}

func (qs *QiitaService) Crawl() {
	qiitaResults := crawlQiita(qs.now, qs.tags)

	// Qiitaのアイテムを保存
	qiita := getQiita(qs.ss)
	saveQiitaItem(qs.now, qiitaResults, qiita, qs.as)
}

func crawlQiita(now time.Time, tags *entities.Tags) []crawler.QiitaResult {
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
	now time.Time,
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
