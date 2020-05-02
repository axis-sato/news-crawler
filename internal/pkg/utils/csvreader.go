package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	sitesCSVPath = "./internal/pkg/db/seeds/master/sites.csv"
	tagsCSVPath  = "./internal/pkg/db/seeds/master/tags.csv"
)

type site struct {
	ID   int
	Name string
	URL  string
}

func (c site) String() string {
	return fmt.Sprintf("ID: %v, Name: %v, URL: %v", c.ID, c.Name, c.URL)
}

func ReadSitesCSV() ([]site, error) {
	var s []site
	file, err := os.Open(sitesCSVPath)
	if err != nil {
		return s, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	reader := csv.NewReader(file)

	var line []string
	for i := 0; ; i++ {
		line, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		if i == 0 {
			// タイトル行はスキップ
			continue
		}

		var site site
		for i, v := range line {
			switch i {
			case 0:
				site.ID, err = strconv.Atoi(v)
				if err != nil {
					return s, err
				}
			case 1:
				site.Name = v
			case 2:
				site.URL = v
			}
		}

		s = append(s, site)
	}

	return s, nil
}

type tag struct {
	ID         int
	Name       string
	IsFollowed int
}

func (t tag) String() string {
	return fmt.Sprintf("ID: %v, Name: %v, IsFollowed: %v", t.ID, t.Name, t.IsFollowed)
}

func ReadTagsCSV() ([]tag, error) {
	var t []tag
	file, err := os.Open(tagsCSVPath)
	if err != nil {
		return t, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	reader := csv.NewReader(file)

	var line []string
	for i := 0; ; i++ {
		line, err = reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		if i == 0 {
			// タイトル行はスキップ
			continue
		}

		var tag tag
		for i, v := range line {
			switch i {
			case 0:
				tag.ID, err = strconv.Atoi(v)
				if err != nil {
					return t, err
				}
			case 1:
				tag.Name = v
			case 2:
				tag.IsFollowed, err = strconv.Atoi(v)
				if err != nil {
					return t, err
				}
			}
		}

		t = append(t, tag)
	}

	return t, nil
}
