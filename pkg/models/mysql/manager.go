package mysql

import (
	"database/sql"
	"errors"
)

var ErrNoRecord = errors.New(" models: no matching record found")

func Insert(stmt string, db *sql.DB, arg ...interface{}) (int, error) {
	result, err := db.Exec(stmt, arg...)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func Get(stmt string, param interface{}, db *sql.DB, arg ...interface{}) error {
	err := db.QueryRow(stmt, param).Scan(arg...)
	if err == sql.ErrNoRows {
		return ErrNoRecord
	}
	return err
}

func Update(stmt string, db *sql.DB, arg ...interface{}) error {
	_, err := db.Exec(stmt, arg...)

	if err == sql.ErrNoRows {
		return ErrNoRecord
	}
	return err
}
