package main

import (
	"fmt"
	"os"
	"time"

	"github.com/c8112002/news-crawler/crawler"

	"github.com/c8112002/news-crawler/db"
	"github.com/c8112002/news-crawler/store"
	"github.com/c8112002/news-crawler/utils"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := utils.LoadEnv(); err != nil {
		panic(err.Error())
	}

	d, err := db.New(utils.GetEnv())

	if err != nil {
		panic(err.Error())
	}

	defer d.Close()

	ts := store.NewTagStore(d)
	tags, err := ts.GetFollowingTag()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(tags)

	today := time.Now()
	_1weekAgo := today.AddDate(0, 0, -7)
	ac := crawler.ArticleCrawler{
		Token: os.Getenv("QIITA_TOKEN"),
		Tags:  tags.Names(),
		From:  _1weekAgo,
		To:    today,
	}
	articles, err := ac.Run()
	if err != nil {
		panic(err.Error())
	}

	for _, a := range articles {
		fmt.Println(a)
	}
}
