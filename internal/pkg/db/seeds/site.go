package seeds

import (
	"database/sql"
	"fmt"

	"github.com/c8112002/news-crawler/internal/pkg/utils"
)

func MakeSites(db *sql.DB) {

	sites, _ := utils.ReadSitesCSV()

	for _, site := range sites {
		ins, err := db.Prepare("INSERT INTO sites(id, name, url) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		_, err = ins.Exec(site.ID, site.Name, site.URL)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Printf("sitesテーブルにマスタデータを投入しました。\nマスタデータ: %v\n", sites)
}
