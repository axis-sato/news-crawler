package main

import (
	"fmt"
	"github.com/c8112002/news-crawler/crawler"
	"github.com/c8112002/news-crawler/utils"
	"os"
	"time"
)

func main() {
	if err := utils.LoadEnv(); err != nil {
		panic(err.Error())
	}

	today := time.Now()
	_1weekAgo := today.AddDate(0,0, -7)
	ac := crawler.ArticleCrawler{
		Token: os.Getenv("QIITA_TOKEN"),
		Tags:  []string{"go", "kotlin"},
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
