package server

import (
	"fmt"
	"os"

	"spark/config"
	"spark/handlers/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {

	config := config.GetServerConfig()
	port := os.Getenv("PORT")

	if os.Getenv("MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Gin Default(Logger, Recovery)
	router := gin.Default()

	// Proxy
	router.SetTrustedProxies(config.Proxies)

	// Cors
	router.Use(cors.New(config.Cors))

	// router
	router.GET("/users", users.GetUser)

	router.Run(fmt.Sprintf(":%v", port))

	select {}
}
