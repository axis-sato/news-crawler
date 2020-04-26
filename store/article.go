package store

import (
	"database/sql"

	"github.com/c8112002/news-crawler/db"

	"github.com/c8112002/news-crawler/entities"
)

type ArticleStore struct {
	db *sql.DB
}

func NewArticleStore(db *sql.DB) *ArticleStore {
	return &ArticleStore{db: db}
}

func (as *ArticleStore) SaveArticles(articles []*entities.Article, tag *entities.Tag, site *entities.Site) error {
	if err := db.Transaction(func(tx *sql.Tx) error {
		for _, a := range articles {
			err := as.saveArticle(a, tag, site, tx)
			if err != nil {
				return err
			}
		}
		return nil
	}, as.db); err != nil {
		return err
	}
	return nil
}

func (as *ArticleStore) saveArticle(article *entities.Article, tag *entities.Tag, site *entities.Site, tx *sql.Tx) error {
	articleInsert, err := tx.Prepare("INSERT INTO articles(title, url, image, crawled_at, sites_id, original_id) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	res, err := articleInsert.Exec(article.Title, article.URL, article.URL, article.CrawledAt, site.ID, article.OriginalID)
	if err != nil {
		return err
	}
	lastArticleID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	articleTagsInsert, err := tx.Prepare("INSERT INTO article_tags(article_id, tag_id) VALUES(?,?)")
	if err != nil {
		return err
	}
	_, err = articleTagsInsert.Exec(lastArticleID, tag.ID)
	if err != nil {
		return err
	}

	return nil
}
