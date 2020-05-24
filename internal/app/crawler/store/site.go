package store

import (
	"database/sql"

	"github.com/c8112002/news-crawler/internal/pkg/entities"
)

const (
	qiitaID    = 1
	devDotToID = 2
	hatenaID   = 3
)

type SiteStore struct {
	db *sql.DB
}

func NewSiteStore(db *sql.DB) *SiteStore {
	return &SiteStore{db: db}
}

func (ss *SiteStore) GetSite(siteID int) (*entities.Site, error) {
	var site entities.Site
	err := ss.db.QueryRow("SELECT id, name, url FROM sites WHERE ID = ?", siteID).Scan(&site.ID, &site.Name, &site.URL)
	if err != nil {
		return nil, err
	}

	return &site, nil
}

func (ss *SiteStore) GetQiita() (*entities.Site, error) {
	return ss.GetSite(qiitaID)
}

func (ss *SiteStore) GetDevTo() (*entities.Site, error) {
	return ss.GetSite(devDotToID)
}
