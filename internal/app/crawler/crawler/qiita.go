package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/c8112002/news-crawler/internal/pkg/entities"

	"github.com/PuerkitoBio/goquery"
)

const (
	qiitaBaseURL    = "https://qiita.com/api/v2/"
	qiitaAction     = "items"
	qiitaBaseParams = "?page=1&per_page=100"
	qiitaTimeFormat = "2006-01-02"
)

type QiitaCrawler struct {
	Token string
	Tags  *entities.Tags
	From  time.Time
	To    time.Time
}

func NewQiitaCrawler(token string, tags *entities.Tags, from time.Time, to time.Time) *QiitaCrawler {
	return &QiitaCrawler{
		Token: token,
		Tags:  tags,
		From:  from,
		To:    to,
	}
}

func (qc *QiitaCrawler) Run() ([]CrawlingResult, error) {
	fromDay := qc.From.Format(qiitaTimeFormat)
	toDay := qc.To.Format(qiitaTimeFormat)

	var results []CrawlingResult

	for _, tag := range *qc.Tags {

		query := fmt.Sprintf("&query=tag:%s+created:>=%s+created:<=%s", tag.Name, fromDay, toDay)

		endpoint, err := url.Parse(qiitaBaseURL + qiitaAction + qiitaBaseParams + query)
		if err != nil {
			return results, err
		}

		var header http.Header

		if len(qc.Token) > 0 {
			header = http.Header{
				"Content-Type":  {"application/json"},
				"Authorization": {"Bearer " + qc.Token},
			}
		} else {
			header = http.Header{
				"Content-Type": {"application/json"},
			}
		}

		resp, err := http.DefaultClient.Do(&http.Request{
			Method: http.MethodGet,
			URL:    endpoint,
			Header: header,
		})

		if err != nil {
			panic(err.Error())
		}

		results, err = func() ([]CrawlingResult, error) {
			defer func() {
				if err := resp.Body.Close(); err != nil {
					panic(err)
				}
			}()

			if err != nil {
				return results, err
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return results, err
			}

			var responses []qiitaResponse
			if err := json.Unmarshal(b, &responses); err != nil {
				return results, err
			}

			var resp []ArticleResponse
			for _, item := range extractPopularItems(responses) {
				crawlThumbnail(&item)
				resp = append(resp, item)
			}
			t := tag // tagはループの度に上書きされてしまうのでここでコピーする
			res := CrawlingResult{Tag: &t, Responses: resp}
			results = append(results, res)
			return results, nil
		}()
	}

	return results, nil
}

func extractPopularItems(source []qiitaResponse) []qiitaResponse {
	var articles []qiitaResponse
	for _, a := range source {
		if a.Likes >= 10 {
			articles = append(articles, a)
		}
	}

	return articles
}

func crawlThumbnail(item *qiitaResponse) {
	res, err := http.Get(item.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err.Error())
		}
	}()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	thumbnail, _ := doc.Find("meta[property='og:image']").First().Attr("content")
	item.Thumbnail = thumbnail
}

type qiitaResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Likes     int    `json:"likes_count"`
	Thumbnail string
}

func (r qiitaResponse) ToArticle() *Article {
	return NewArticle(r.ID, r.Title, r.URL, r.Likes, r.Thumbnail)
}
