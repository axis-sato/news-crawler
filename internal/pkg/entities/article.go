package entities

import (
	"fmt"
	"time"
)

type Article struct {
	ID         int64
	OriginalID string
	Title      string
	URL        string
	Image      string
	CrawledAt  time.Time
	Site       *Site
	Tags       *Tags
}

func NewArticle(originalID string, title string, url string, image string, crawledAt time.Time, site *Site, tags *Tags) *Article {
	return &Article{
		OriginalID: originalID,
		Title:      title,
		URL:        url,
		Image:      image,
		CrawledAt:  crawledAt,
		Site:       site,
		Tags:       tags,
	}
}

func (a *Article) IsDuplicatedWith(target *Article) bool {
	return a.Site.ID == target.Site.ID &&
		a.OriginalID == target.OriginalID
}

func (a Article) String() string {
	return fmt.Sprintf("{id: %v, title: %v, url: %v, image: %v, crawledAt: %v, site: %v, tags: %v}", a.ID, a.Title, a.URL, a.Image, a.CrawledAt, a.Site, a.Tags)
}
