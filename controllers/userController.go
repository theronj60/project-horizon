package controllers

import (
	// "github.com/theronj60/project-horizon/db"
	"github.com/gin-gonic/gin"
	"github.com/theronj60/project-horizon/initializers"
	"github.com/theronj60/project-horizon/models"
)

type Message struct {
	id      int
	Message string
}

func UserIndex(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
	// returns all users
}
func UserShow(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	initializers.DB.First(&user, id)
	c.JSON(200, gin.H{
		"user": user,
	})
	// returns one user
}
func create() {
	// creates a user
	// return message
}
func update() {
	// updates a user
	// return message
}
func delete() {
	// deletes a user
	// return message
}
