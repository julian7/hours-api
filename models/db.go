package models

import (
	"database/sql"

	// postgres
	_ "github.com/lib/pq"
)

var conn *sql.DB

//InitDB initializes persistent database connection
func InitDB(source string) (*sql.DB, error) {
	var err error
	conn, err = sql.Open("postgres", source)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}
