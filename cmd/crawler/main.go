package main

import (
	"os"
	"time"

	"github.com/c8112002/news-crawler/internal/app/crawler/services"

	log "github.com/sirupsen/logrus"

	"github.com/c8112002/news-crawler/internal/app/crawler/store"
	"github.com/c8112002/news-crawler/internal/pkg/db"
	"github.com/c8112002/news-crawler/internal/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

var now time.Time

func init() {
	now = time.Now()
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}

func main() {
	initEnv()

	d, err := db.New()
	if err != nil {
		panic(err.Error())
	}
	defer d.Close()

	ts := store.NewTagStore(d)
	ss := store.NewSiteStore(d)
	as := store.NewArticleStore(d)

	// クロール対象のタグを取得
	tags := services.GetTargetTags(ts)
	log.WithFields(log.Fields{
		"tags": tags,
	}).Debug("クロール対象のタグ")

	// Qiitaのクローリング
	_ = services.NewQiitaService(tags, ss, as).
		Crawl(now)
	// Dev.toのクローリング
	_ = services.NewDevToService(tags, ss, as).
		Crawl(now)

}

func initEnv() {
	if err := utils.Init(); err != nil {
		panic(err.Error())
	}
}
