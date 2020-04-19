package main

import (
	"encoding/json"
	"fmt"
	"github.com/c8112002/news-crawler/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	if err := utils.LoadEnv(); err != nil {
		panic(err.Error())
	}

	today := time.Now()
	_1weekAgo := today.AddDate(0,0, -7)
	articles, err := fetchDataFromQiita(os.Getenv("QIITA_TOKEN"), "go", _1weekAgo, today)
	if err != nil {
		panic(err.Error())
	}

	for _, a := range articles {
		fmt.Println(a)
	}
}

type Article struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Likes int    `json:"likes_count"`
}

func (a Article) String() string {
	return fmt.Sprintf("id: %v, title: %v, url: %v, likes: %v", a.ID, a.Title, a.URL, a.Likes)
}

func fetchDataFromQiita(token string, tag string, from time.Time, to time.Time) ([]Article, error) {
	baseUrl := "https://qiita.com/api/v2/"
	action := "items"
	baseParams := "?page=1&per_page=100"
	fromDay := from.Format("2006-01-02")
	toDay := to.Format("2006-01-02")
	query := fmt.Sprintf("&query=tag:%s+created:>=%s+created:<=%s", tag, fromDay, toDay)

	var articles []Article

	endpoint, err := url.Parse(baseUrl + action + baseParams + query)
	if err != nil {
		return articles, err
	}

	var header http.Header

	if len(token) > 0 {
		header = http.Header{
			"Content-Type": {"application/json"},
			"Authorization": {"Bearer " + token},
		}
	} else {
		header = http.Header{
			"Content-Type": {"application/json"},
		}
	}

	resp, err := http.DefaultClient.Do(&http.Request{
		Method:           http.MethodGet,
		URL:              endpoint,
		Header: header,
	})

	defer func() {
		if err := resp.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		return articles, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return articles, err
	}

	if err := json.Unmarshal(b, &articles); err != nil {
		return articles, err
	}

	return extractPopularArticles(articles), nil
}

func extractPopularArticles(source []Article) []Article {
	var articles []Article
	for _, a := range source {
		if a.Likes >= 10 {
			articles = append(articles, a)
		}
	}

	return articles
}