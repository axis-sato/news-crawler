package services

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/c8112002/news-crawler/internal/app/crawler/crawler"
	"github.com/c8112002/news-crawler/internal/app/crawler/store"
	"github.com/c8112002/news-crawler/internal/pkg/entities"
)

func saveArticle(
	now time.Time,
	results []crawler.CrawlingResult,
	site *entities.Site,
	as *store.ArticleStore,
	ss *store.SiteStore,
) error {

	log.WithFields(log.Fields{
		"article": results,
	}).Debug(fmt.Sprintf("取得した記事一覧"))

	for _, r := range results {

		storedArticles, _ := as.GetAllArticlesTaggedAs(r.Tag, ss)

		var articles []*entities.Article
		for _, res := range r.Responses {
			article := res.ToArticle()
			tag := entities.NewTag(r.Tag.ID, r.Tag.Name)
			tags := entities.NewTags([]entities.Tag{*tag})
			a := entities.NewArticle(article.ID, article.Title, article.URL, article.Thumbnail, now, site, tags)
			if containsDuplicatedArticle(a, storedArticles) {
				log.WithFields(log.Fields{
					"article": a,
				}).Debug(fmt.Sprintf("既に保存済みなのでスキップした記事のタイトル: %v", a.Title))
			} else {
				articles = append(articles, a)
			}
		}

		res, err := as.SaveArticles(articles)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		log.WithFields(log.Fields{
			"articles": res,
		}).Debug(fmt.Sprintf("保存した%vのアイテム", site.Name))
	}

	return nil
}

func containsDuplicatedArticle(target *entities.Article, sources []*entities.Article) bool {
	for _, a := range sources {
		if target.IsDuplicatedWith(a) {
			return true
		}
	}
	return false
}
