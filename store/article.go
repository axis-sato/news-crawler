package store

import (
	"context"
	"database/sql"

	"github.com/c8112002/news-crawler/entities"
)

type ArticleStore struct {
	db *sql.DB
}

func NewArticleStore(db *sql.DB) *ArticleStore {
	return &ArticleStore{db: db}
}

func (as *ArticleStore) SaveArticles(articles []*entities.Article, tag *entities.Tag, site *entities.Site) error {
	ctx := context.Background()
	tx, err := as.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	for _, a := range articles {
		err := as.saveArticle(tx, a, tag, site)
		if err != nil {
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (as *ArticleStore) saveArticle(tx *sql.Tx, article *entities.Article, tag *entities.Tag, site *entities.Site) error {
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
