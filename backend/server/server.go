package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"storage/database"
)

// initialized the Server and creates the first table products
func InitializeServer() error {
	r := gin.Default()
	r.GET("/new", func(c *gin.Context) {
		database.InitDatabase()
		c.JSON(200, gin.H{
			"message": "Table created successfully",
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080
	return nil
}
