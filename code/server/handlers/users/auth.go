package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	response := Response{
		Ok: true,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func SignUp(c *gin.Context) {

	response := Response{
		Ok: true,
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
