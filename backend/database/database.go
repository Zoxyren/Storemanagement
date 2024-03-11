package database

import (
	"database/sql"
	"fmt"
	"log"
)

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
		log.Fatal("Error opening database", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products(ID INTEGER PRIMARY KEY, DESCRIPTION VARCHAR, PRICE INTEGER, AMOUNT FLOAT)")
	if err != nil {
		log.Fatal("Failed to create table:", err)

	}
	defer db.Close()
	return db, err

}
