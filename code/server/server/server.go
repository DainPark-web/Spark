package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	port := 8080

	router := gin.Default()

	router.Run(fmt.Sprintf(":%d", port))

	select {}
}
