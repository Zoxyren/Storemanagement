package database

import (
	"database/sql"
	"fmt"
	"log/slog"
)

// Todo: add .env variable for database
const (
	host     = "localhost"
	port     = 5433
	user     = "admin"
	password = "123321"
	dbname   = "storage_db"
)

func InitDatabase() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		slog.Error("Error opening database", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products(ID INTEGER PRIMARY KEY, NAME VARCHAR,  DESCRIPTION VARCHAR, PRICE INTEGER, AMOUNT FLOAT)")
	if err != nil {
		slog.Error("Failed to create table:", err)

	}
	return db, err

}
