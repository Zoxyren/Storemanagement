package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"storage/server"
)

func main() {
	// Gin-Router initialisieren
	r := gin.Default()
	// Routen definieren
	//server.DefineRoutes(r)
	server.InitializeServer()
	// Server starten
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
