package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// serve react index.html
	router.Use(static.Serve("/", static.LocalFile("../frontend/dist", true)))

	// setup api route
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})
	}

	router.Run(":3000")
}
