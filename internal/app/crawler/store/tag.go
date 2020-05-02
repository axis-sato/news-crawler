package store

import (
	"database/sql"

	"github.com/c8112002/news-crawler/internal/pkg/entities"
)

type TagStore struct {
	db *sql.DB
}

func NewTagStore(db *sql.DB) *TagStore {
	return &TagStore{db: db}
}

func (ts *TagStore) GetFollowingTag() (*entities.Tags, error) {
	rows, err := ts.db.Query("SELECT id, name FROM tags WHERE is_followed = 1")
	if err != nil {
		return nil, err
	}

	var tags entities.Tags
	for rows.Next() {
		tag := entities.Tag{}
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return &tags, nil
}
