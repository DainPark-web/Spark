package users

import (
	"net/http"

	"spark/database"
	"spark/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	database.DB.Find(&users)
	if len(users) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"users": users})
}

func GetUser(c *gin.Context) {
	user := models.User{}
	result := database.DB.Select("name").Where("name = ?", "John Doe").Take(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}
