package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := "host=localhost port=5432 user=user password=password dbname=datasystem sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to PostgreSQL successfully")
	return nil
}

func InsertData(source string, value float64) error {
	_, err := DB.Exec("INSERT INTO collected_data (source, value) VALUES ($1, $2)", source, value)
	return err
}
