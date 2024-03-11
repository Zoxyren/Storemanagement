package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "net/http"
	"storage/controller"
	"storage/database"
)

// InitializeServer initialized the Server and creates the first table products
func InitializeServer() error {
	r := gin.Default()

	// GET-Route
	r.GET("/new", func(c *gin.Context) {
		database.InitDatabase()
		c.JSON(200, gin.H{
			"message": "Tabelle erfolgreich erstellt",
		})
	})

	// POST-Route
	r.POST("/product", func(c *gin.Context) {
		err := controller.InsertProduct(c.Writer, c.Request)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to insert product",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Produkt erfolgreich erstellt",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
	return nil
}
