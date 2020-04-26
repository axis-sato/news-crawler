package db

import (
	"database/sql"
	"log"

	"github.com/c8112002/news-crawler/utils"
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

func Transaction(txFunc func(*sql.Tx) error, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			log.Println("Recover")
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println("Rollback")
			err = tx.Rollback()
		} else {
			log.Println("Commit")
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)
	return err
}
