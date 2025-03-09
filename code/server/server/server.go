package server

import (
	"fmt"
	"os"

	"spark/config"
	"spark/routes"

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
	r := gin.Default()

	// Proxy
	r.SetTrustedProxies(config.Proxies)

	// Cors
	r.Use(cors.New(config.Cors))

	// routers
	routes.SetupRoutes(r)

	r.Run(fmt.Sprintf(":%v", port))

}
