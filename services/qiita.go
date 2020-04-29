package services

import (
	"os"
	"time"

	"github.com/c8112002/news-crawler/crawler"
	"github.com/c8112002/news-crawler/entities"
	"github.com/c8112002/news-crawler/store"
)

type QiitaService struct {
	tags *entities.Tags
	ss   *store.SiteStore
	as   *store.ArticleStore
}

func NewQiitaService(tags *entities.Tags, ss *store.SiteStore, as *store.ArticleStore) *QiitaService {
	return &QiitaService{tags: tags, ss: ss, as: as}
}

func (qs *QiitaService) Crawl(now time.Time) error {
	results := crawlQiita(now, qs.tags)

	// Qiitaのアイテムを保存
	qiita, err := qs.ss.GetQiita()
	if err != nil {
		return err
	}
	return saveArticle(now, results, qiita, qs.as)
}

func crawlQiita(now time.Time, tags *entities.Tags) []crawler.CrawlingResult {
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
