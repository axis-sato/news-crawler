package crawler

import (
	"fmt"

	"github.com/c8112002/news-crawler/internal/pkg/entities"
)

type ArticleResponse interface {
	ToArticle() *Article
}

type Article struct {
	ID        string
	Title     string
	URL       string
	Likes     int
	Thumbnail string
}

func NewArticle(id string, title string, url string, like int, thumbnail string) *Article {
	return &Article{
		ID:        id,
		Title:     title,
		URL:       url,
		Likes:     like,
		Thumbnail: thumbnail,
	}
}

func (a *Article) String() string {
	return fmt.Sprintf("{id: %v, title: %v, url: %v, likes: %v, thumbnail: %v}", a.ID, a.Title, a.URL, a.Likes, a.Thumbnail)
}

type CrawlingResult struct {
	Tag       *entities.Tag
	Responses []ArticleResponse
}
