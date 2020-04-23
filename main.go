package main

import (
	"fmt"
	"os"
	"time"

	"github.com/c8112002/news-crawler/entities"

	"github.com/c8112002/news-crawler/crawler"

	"github.com/c8112002/news-crawler/db"
	"github.com/c8112002/news-crawler/store"
	"github.com/c8112002/news-crawler/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	loadEnv()

	d, err := db.New(utils.GetEnv())

	if err != nil {
		panic(err.Error())
	}

	defer d.Close()

	ts := store.NewTagStore(d)

	tags := getTags(ts)

	today := time.Now()
	_1weekAgo := today.AddDate(0, 0, -7)
	ac := crawler.NewArticleCrawler(
		os.Getenv("QIITA_TOKEN"),
		tags.Names(),
		_1weekAgo,
		today,
	)
	articles, err := ac.Run()
	if err != nil {
		panic(err.Error())
	}

	for _, a := range articles {
		fmt.Println(a)
	}
}

func loadEnv() {
	if err := utils.LoadEnv(); err != nil {
		panic(err.Error())
	}
}

func getTags(ts *store.TagStore) *entities.Tags {
	tags, err := ts.GetFollowingTag()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(tags)

	return tags
}
