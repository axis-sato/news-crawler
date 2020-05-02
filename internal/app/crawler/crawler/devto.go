package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/c8112002/news-crawler/internal/pkg/entities"
)

const (
	devToBaseURL    = "https://dev.to/api/"
	devToAction     = "articles"
	devToBaseParams = "?page=1&per_page=5"
)

type DevToCrawler struct {
	Tags      *entities.Tags
	LastNDays int
}

func NewDevToCrawler(tags *entities.Tags, lastNDays int) *DevToCrawler {
	return &DevToCrawler{
		Tags:      tags,
		LastNDays: lastNDays,
	}
}

func (dc *DevToCrawler) Run() ([]CrawlingResult, error) {
	var results []CrawlingResult
	for _, tag := range *dc.Tags {

		query := fmt.Sprintf("%v&tag=%v&top=%v", devToBaseParams, tag.Name, dc.LastNDays)

		endpoint, err := url.Parse(devToBaseURL + devToAction + devToBaseParams + query)
		if err != nil {
			return results, err
		}

		resp, err := http.DefaultClient.Do(&http.Request{
			Method: http.MethodGet,
			URL:    endpoint,
		})
		if err != nil {
			return results, err
		}

		results, err = func() ([]CrawlingResult, error) {
			defer func() {
				if err := resp.Body.Close(); err != nil {
					panic(err.Error())
				}
			}()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return results, err
			}

			var responses []devToResponse

			if err := json.Unmarshal(b, &responses); err != nil {
				return results, err
			}

			var resp []ArticleResponse
			for _, dr := range responses {
				resp = append(resp, dr)
			}

			t := tag // tagはループの度に上書きされてしまうのでここでコピーする
			res := CrawlingResult{
				Tag:       &t,
				Responses: resp,
			}
			results = append(results, res)

			return results, nil

		}()

		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

type devToResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Likes     int    `json:"positive_reactions_count"`
	Thumbnail string `json:"social_image"`
}

func (r devToResponse) ToArticle() *Article {
	return NewArticle(strconv.Itoa(r.ID), r.Title, r.URL, r.Likes, r.Thumbnail)
}
