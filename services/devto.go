package services

import (
	"github.com/c8112002/news-crawler/crawler"
	"github.com/c8112002/news-crawler/entities"
	log "github.com/sirupsen/logrus"
)

const lastNDays = 7

type DevToService struct {
	tags *entities.Tags
}

func NewDevToService(tags *entities.Tags) *DevToService {
	return &DevToService{tags: tags}
}

func (ds *DevToService) Crawl() {
	results := crawlDevTo(ds.tags, lastNDays)
	log.WithFields(log.Fields{
		"results": results,
	}).Debug("Dev.toの取得記事")
}

func crawlDevTo(tags *entities.Tags, lastNDays int) []crawler.DevToResult {
	results, err := crawler.NewDevToCrawler(tags, lastNDays).
		Run()
	if err != nil {
		panic(err.Error())
	}
	return results
}
