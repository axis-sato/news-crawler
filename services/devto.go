package services

import (
	"time"

	"github.com/c8112002/news-crawler/crawler"
	"github.com/c8112002/news-crawler/entities"
	"github.com/c8112002/news-crawler/store"
	log "github.com/sirupsen/logrus"
)

const lastNDays = 7

type DevToService struct {
	tags *entities.Tags
	ss   *store.SiteStore
	as   *store.ArticleStore
}

func NewDevToService(tags *entities.Tags, ss *store.SiteStore, as *store.ArticleStore) *DevToService {
	return &DevToService{tags: tags, ss: ss, as: as}
}

func (ds *DevToService) Crawl(now time.Time) error {
	results := crawlDevTo(ds.tags, lastNDays)
	log.WithFields(log.Fields{
		"results": results,
	}).Debug("Dev.toの取得記事")

	// Qiitaのアイテムを保存
	devto, err := ds.ss.GetDevTo()
	if err != nil {
		return err
	}
	return saveArticle(now, results, devto, ds.as)
}

func crawlDevTo(tags *entities.Tags, lastNDays int) []crawler.CrawlingResult {
	results, err := crawler.NewDevToCrawler(tags, lastNDays).
		Run()
	if err != nil {
		panic(err.Error())
	}
	return results
}
