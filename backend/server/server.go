package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log/slog"
	_ "net/http"
	"storage/controller"
	"storage/database"
)

// InitializeServer initialized the Server and creates the first table products
func InitializeServer() error {
	r := gin.Default()

	// GET-Route
	r.GET("/new", func(c *gin.Context) {
		_, err := database.InitDatabase()
		if err != nil {
			return
		}
		c.JSON(200, gin.H{
			"message": "Tabelle erfolgreich erstellt",
		})
	})

	// POST-Route
	r.POST("/product", func(c *gin.Context) {
		database.InitDatabase()
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
	r.GET("/products", func(c *gin.Context) {
		database.InitDatabase()
		_, err := controller.GetProducts(c.Writer, c.Request)
		if err != nil {
			c.JSON(201, gin.H{
				"message": "Produkt erfolgreich geladen",
			})
		}
		return
	})

	err := r.Run()
	if err != nil {
		slog.Error("Failed to start webserver", err)
	} // listen and serve on 0.0.0.0:8080
	return nil
}
