package services

import (
	"fmt"
	"time"

	"github.com/c8112002/news-crawler/crawler"
	"github.com/c8112002/news-crawler/entities"
	"github.com/c8112002/news-crawler/store"
	log "github.com/sirupsen/logrus"
)

func saveArticle(
	now time.Time,
	results []crawler.CrawlingResult,
	site *entities.Site,
	as *store.ArticleStore,
) error {
	for _, r := range results {
		var articles []*entities.Article
		for _, res := range r.Responses {
			article := res.ToArticle()
			a := entities.NewArticle(article.ID, article.Title, article.URL, article.Thumbnail, now)
			articles = append(articles, a)
		}
		res, err := as.SaveArticles(articles, r.Tag, site)
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{
			"articles": res,
		}).Debug(fmt.Sprintf("保存した%vのアイテム", site.Name))
	}

	return nil
}
