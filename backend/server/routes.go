package server

import (
	"github.com/gin-gonic/gin"
	"storage/database"
)

func DefineRoutes(r *gin.Engine) {
	router := gin.Default()
	router.GET("/table", func(c *gin.Context) {
		database.CreateTable()
		c.JSON(200, gin.H{
			"message": "Table has been created successfully",
		})

	})
}
