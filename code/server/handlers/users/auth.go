package users

import (
	"net/http"

	"spark/database"
	"spark/models"

	"spark/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	response := Response{
		Ok: true,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func SignUp(c *gin.Context) {
	newUser := models.User{
		Email:    c.PostForm("email"),
		Name:     c.PostForm("name"),
		Password: services.HashPassword(c.PostForm("password")),
	}

	if newUser.Name == "" || newUser.Password == "" || newUser.Email == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Name, password and email are required"})
		return
	}

	existingUser := models.User{}
	database.DB.Where("email = ?", newUser.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User with this email already exists"})
		return
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	response := SignUpResponse{
		Ok: true,
		Data: gin.H{
			"user": newUser,
		},
	}

	c.IndentedJSON(http.StatusOK, response)
}

func Logout(c *gin.Context) {

	response := Response{
		Ok: true,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func Getme(c *gin.Context) {
	response := Response{
		Ok: true,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {

	response := Response{
		Ok: true,
	}

	c.IndentedJSON(http.StatusOK, response)

}

func DeleteUser(c *gin.Context) {

	response := Response{
		Ok: true,
	}

	c.IndentedJSON(http.StatusOK, response)
}
