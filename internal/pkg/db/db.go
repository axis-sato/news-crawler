package db

import (
	"database/sql"
)

func New() (*sql.DB, error) {
	c, err := readDBConf()

	if err != nil {
		return nil, err
	}

	return sql.Open(c.Dialect, c.Datasource)
}

func Transaction(txFunc func(*sql.Tx) error, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			err = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = txFunc(tx)
	return err
}
