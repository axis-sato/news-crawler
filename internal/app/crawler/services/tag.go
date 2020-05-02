package services

import (
	"github.com/c8112002/news-crawler/internal/app/crawler/store"
	"github.com/c8112002/news-crawler/internal/pkg/entities"
)

func GetTargetTags(ts *store.TagStore) *entities.Tags {
	tags, err := ts.GetFollowingTag()
	if err != nil {
		panic(err.Error())
	}

	return tags
}
