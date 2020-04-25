package main

import (
	"fmt"
	"os"
	"time"

	"github.com/c8112002/news-crawler/crawler"

	"github.com/c8112002/news-crawler/entities"

	"github.com/c8112002/news-crawler/db"
	"github.com/c8112002/news-crawler/store"
	"github.com/c8112002/news-crawler/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	now := time.Now()

	loadEnv()

	d, err := db.New(utils.GetEnv())

	if err != nil {
		panic(err.Error())
	}

	defer d.Close()

	ts := store.NewTagStore(d)
	ss := store.NewSiteStore(d)
	as := store.NewArticleStore(d)

	tags := getTags(ts)

	today := time.Now()
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

	qiita := getQiita(ss)
	for _, r := range qiitaResults {
		var articles []*entities.Article
		for _, item := range r.Items {
			a := entities.NewArticle(item.ID, item.Title, item.URL, item.Thumbnail, now)
			articles = append(articles, a)
		}
		_ = as.SaveArticles(articles, r.Tag, qiita)
	}
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

	fmt.Println(tags)

	return tags
}

func getQiita(ss *store.SiteStore) *entities.Site {
	qiita, err := ss.GetQiita()
	if err != nil {
		panic(err.Error())
	}

	return qiita
}
