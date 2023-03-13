package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/theronj60/project-horizon/controllers"
	"github.com/theronj60/project-horizon/initializers"
)

func init() {
	initializers.Connect()
}

func main() {
	router := gin.Default()

	// serve react index.html
	router.Use(static.Serve("/", static.LocalFile("./views/dist", true)))

	// setup api route
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World",
			})
		})
		api.GET("/users", controllers.UserIndex)
		api.GET("/user/:id", controllers.UserShow)
	}

	router.Run(":3000")
}
