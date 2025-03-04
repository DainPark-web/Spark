package server

import (
	"fmt"
	"os"

	"spark/handlers/users"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	port := os.Getenv("PORT")

	router := gin.Default()

	router.GET("/users", users.GetUser)

	router.Run(fmt.Sprintf(":%v", port))

	select {}
}
