package routes

import (
	"spark/handlers/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		SetVersionRoutes(api)
	}
}

func SetVersionRoutes(r *gin.RouterGroup) {
	version1 := r.Group("/v1")
	{
		// User Routes
		SetUserRoutes(version1)
	}
}

func SetUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/users")
	{
		user.GET("/", users.GetUsers)
		user.GET("/:id", users.GetUser)
	}
}
