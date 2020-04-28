package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/c8112002/news-crawler/entities"
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

func (dc *DevToCrawler) Run() ([]DevToResult, error) {
	var results []DevToResult
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

		var articles []devToArticle
		results, err = func() ([]DevToResult, error) {
			defer func() {
				if err := resp.Body.Close(); err != nil {
					panic(err.Error())
				}
			}()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return results, err
			}

			if err := json.Unmarshal(b, &articles); err != nil {
				return results, err
			}

			t := tag // tagはループの度に上書きされてしまうのでここでコピーする
			res := DevToResult{
				Tag:      &t,
				Articles: articles,
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

type devToArticle struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Likes     int    `json:"positive_reactions_count"`
	Thumbnail string `json:"social_image"`
}

func (d *devToArticle) String() string {
	return fmt.Sprintf("{id: %v, title: %v, url: %v, likes: %v, thumbnail: %v}", d.ID, d.Title, d.URL, d.Likes, d.Thumbnail)
}

type DevToResult struct {
	Tag      *entities.Tag
	Articles []devToArticle
}
