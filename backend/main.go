package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"storage/controller"
	"storage/database"
	"storage/server"
)

func main() {
	// Gin-Router initialisieren
	r := gin.Default()
	datadb, err := database.InitDatabase()
	if err != nil {
		slog.Error("Error initializing database", "Message", err)
		return
	}
	defer func(datadb *sql.DB) {
		err := datadb.Close()
		if err != nil {
			slog.Error("Error closing database", "Message", err)
		}
	}(datadb)
	controller.DB = *datadb
	if err != nil {
		slog.Error("Error initializing database", "Message")
	}
	// Routen definieren
	//server.DefineRoutes(r)
	er := server.InitializeServer()
	if err != nil {
		slog.Error("Error initializing server", "Message", er)
		return
	}
	// Server starten
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
