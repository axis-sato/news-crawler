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

func (as *ArticleStore) SaveArticles(articles []*entities.Article, tag *entities.Tag, site *entities.Site) ([]*entities.Article, error) {
	var res []*entities.Article
	err := db.Transaction(func(tx *sql.Tx) error {
		for _, a := range articles {
			article, err := as.saveArticle(a, tag, site, tx)
			if err != nil {
				return err
			}
			res = append(res, article)
		}
		return nil
	}, as.db)

	return res, err
}

func (as *ArticleStore) saveArticle(article *entities.Article, tag *entities.Tag, site *entities.Site, tx *sql.Tx) (*entities.Article, error) {
	articleInsert, err := tx.Prepare("INSERT INTO articles(title, url, image, crawled_at, sites_id, original_id) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	res, err := articleInsert.Exec(article.Title, article.URL, article.URL, article.CrawledAt, site.ID, article.OriginalID)
	if err != nil {
		return nil, err
	}
	lastArticleID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	articleTagsInsert, err := tx.Prepare("INSERT INTO article_tags(article_id, tag_id) VALUES(?,?)")
	if err != nil {
		return nil, err
	}
	_, err = articleTagsInsert.Exec(lastArticleID, tag.ID)
	if err != nil {
		return nil, err
	}

	article.ID = lastArticleID
	article.Site = site
	article.Tags = &entities.Tags{*tag}

	return article, nil
}
