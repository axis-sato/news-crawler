package seeds

import (
	"database/sql"
	"fmt"

	"github.com/c8112002/news-crawler/internal/pkg/utils"
)

func MakeTags(db *sql.DB) {

	tags, _ := utils.ReadTagsCSV()

	for _, tag := range tags {
		ins, err := db.Prepare("INSERT INTO tags(id, name, is_followed) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		_, err = ins.Exec(tag.ID, tag.Name, tag.IsFollowed)
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Printf("tagsテーブルにマスタデータを投入しました。\nマスタデータ: %v\n", tags)
}
