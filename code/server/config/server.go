package config

import (
	"os"

	"github.com/gin-contrib/cors"
)

type ServerConfig struct {
	Cors    cors.Config
	Proxies []string
}

func GetServerConfig() ServerConfig {
	env := os.Getenv("ENV")
	if env == "dev" {
		return ServerConfig{
			Cors: cors.Config{
				AllowOrigins:     []string{"http://localhost:3000"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
				AllowCredentials: true,
			},
			Proxies: []string{"127.0.0.1"},
		}
	}
	return ServerConfig{
		Cors: cors.Config{
			AllowOrigins:     []string{"https://spark.com"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		},
		Proxies: []string{"your-production-proxy-ip"},
	}
}
