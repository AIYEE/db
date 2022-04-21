package db

import "errors"

var (
	ErrNoDriverDb = errors.New("no database driver")
)

func New(driver string, dbFile string) (DbOperator, error) {
	switch driver {
	case "sqlite3":
		return NewSqlite3(dbFile)
	default:
		return nil, ErrNoDriverDb
	}
}
