package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const dbConnStr = "user=postgres dbname=bank sslmode=disable password=postgres"

func NewConnection() (*sql.DB, error) {
	return sql.Open("postgres", dbConnStr)
}
