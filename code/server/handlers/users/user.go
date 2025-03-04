package users

import (
	"net/http"

	"spark/database"
	"spark/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := models.User{}
	result := database.DB.Select("name").Where("name = ?", "John Doe").Take(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}
