package routes

import (
	"example/userProject/pkg/userData/handlers"

	"github.com/gin-gonic/gin"
)

// Routes function to serve endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{

		v1.POST("user/create", handlers.CreateUser)
		v1.GET("user", handlers.GetAllUsers)
		v1.GET("user/:id", handlers.GetOneUser)
		v1.DELETE("user/:id", handlers.DeleteUser)
		v1.PUT("user/:id", handlers.UpdateUser)

	}

	return router
}
