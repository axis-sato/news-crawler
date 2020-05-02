package db

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func readDBConf() (*dbconf, error) {
	var c dbconf

	viper.SetConfigName("dbconf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/pkg/db")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "c8112002", "news-crawler", "internal", "pkg", "db"))

	if err := viper.ReadInConfig(); err != nil {
		return &c, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return &c, err
	}

	return &c, nil
}

type dbconf struct {
	Development       param
	DevelopmentDocker param
	Production        param
}

type param struct {
	Dialect    string
	Datasource string
	Dir        string
}
