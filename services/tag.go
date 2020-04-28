package services

import (
	"github.com/c8112002/news-crawler/entities"
	"github.com/c8112002/news-crawler/store"
)

func GetTargetTags(ts *store.TagStore) *entities.Tags {
	tags, err := ts.GetFollowingTag()
	if err != nil {
		panic(err.Error())
	}

	return tags
}
