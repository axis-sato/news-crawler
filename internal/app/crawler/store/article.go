package store

import (
	"database/sql"

	"github.com/c8112002/news-crawler/internal/pkg/db"

	"github.com/c8112002/news-crawler/internal/pkg/entities"
)

type ArticleStore struct {
	db *sql.DB
}

func NewArticleStore(db *sql.DB) *ArticleStore {
	return &ArticleStore{db: db}
}

func (as *ArticleStore) SaveArticles(articles []*entities.Article) ([]*entities.Article, error) {
	var res []*entities.Article
	err := db.Transaction(func(tx *sql.Tx) error {
		for _, a := range articles {
			article, err := as.saveArticle(a, tx)
			if err != nil {
				return err
			}
			res = append(res, article)
		}
		return nil
	}, as.db)

	return res, err
}

func (as *ArticleStore) saveArticle(article *entities.Article, tx *sql.Tx) (*entities.Article, error) {
	articleInsert, err := tx.Prepare("INSERT INTO articles(title, url, image, crawled_at, site_id, original_id) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	var res sql.Result
	if article.Image == "" {
		res, err = articleInsert.Exec(article.Title, article.URL, nil, article.CrawledAt, article.Site.ID, article.OriginalID)
	} else {
		res, err = articleInsert.Exec(article.Title, article.URL, article.Image, article.CrawledAt, article.Site.ID, article.OriginalID)
	}
	if err != nil {
		return nil, err
	}
	lastArticleID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	for _, t := range *article.Tags {
		articleTagsInsert, err := tx.Prepare("INSERT INTO article_tags(article_id, tag_id) VALUES(?,?)")
		if err != nil {
			return nil, err
		}
		_, err = articleTagsInsert.Exec(lastArticleID, t.ID)
		if err != nil {
			return nil, err
		}
	}

	article.ID = lastArticleID

	return article, nil
}

func (as *ArticleStore) GetAllArticlesTaggedAs(tag *entities.Tag, ss *SiteStore) ([]*entities.Article, error) {

	rows, err := as.db.Query("SELECT id, title, url, image, crawled_at, site_id, original_id FROM articles INNER JOIN article_tags ON articles.id = article_tags.article_id AND tag_id = ?", tag.ID)
	if err != nil {
		return nil, err
	}

	var articles []*entities.Article
	for rows.Next() {
		article := entities.Article{}
		var siteID int
		if err := rows.Scan(&article.ID, &article.Title, &article.URL, &article.Image, &article.CrawledAt, &siteID, &article.OriginalID); err != nil {
			return nil, err
		}
		site, err := ss.GetSite(siteID)
		if err != nil {
			return nil, err
		}
		article.Site = site

		articles = append(articles, &article)
	}

	return articles, nil
}
