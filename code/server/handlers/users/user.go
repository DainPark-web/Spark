package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	response := Response{
		Ok: true,
	}

	c.IndentedJSON(http.StatusOK, response)
}
