package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ok bool `json:"ok"`
}

func TrackIp(c *gin.Context) {
	response := Response{
		Ok: true,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func StartServer() {
	port := 8080

	router := gin.Default()

	router.GET("/", TrackIp)

	router.Run(fmt.Sprintf(":%d", port))

	select {}
}
