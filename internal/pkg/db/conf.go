package db

import (
	"github.com/c8112002/news-crawler/internal/pkg/utils"
	"github.com/spf13/viper"
	"os"
)

func readDBConf() (*dbconf, error) {
	var c dbconf

	viper.SetConfigName("dbconf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return &c, err
	}

	var key string
	switch utils.AppEnv {
	case utils.Development:
		key = "development"
	case utils.DevelopmentDocker:
		key = "development.docker"
	case utils.Production:
		key = "production"
	default:
		key = "production"
	}

	return &dbconf{
		Dialect:    viper.GetString(key + ".dialect"),
		Datasource: os.ExpandEnv(viper.GetString(key + ".datasource")),
		Dir:        viper.GetString(key + ".dir"),
	}, nil
}

type dbconf struct {
	Dialect    string
	Datasource string
	Dir        string
}
