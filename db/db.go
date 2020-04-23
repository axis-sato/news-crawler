package db

import (
	"database/sql"

	"github.com/c8112002/news-crawler/utils"
	"github.com/spf13/viper"
)

func New(env utils.Env) (*sql.DB, error) {
	c, err := readDBConf()

	if err != nil {
		return nil, err
	}

	switch env {
	case utils.Development:
		return sql.Open(c.Development.Dialect, c.Development.Datasource)
	case utils.DevelopmentDocker:
		return sql.Open(c.DevelopmentDocker.Dialect, c.DevelopmentDocker.Datasource)
	case utils.Production:
		return sql.Open(c.Production.Dialect, c.Production.Datasource)
	default:
		return sql.Open(c.Development.Dialect, c.Development.Datasource)
	}

}

func readDBConf() (*dbconf, error) {
	var c dbconf

	viper.SetConfigName("dbconf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("./db")
	viper.AddConfigPath("../db")

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
