package routes

import (
	handler "example/userProject/pkg/userAuth/handlers"
	"example/userProject/pkg/userData/handlers"

	"example/userProject/pkg/userAuth/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes function to serve endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	public := router.Group("/api")
	public.POST("/register", handler.Signup)
	public.POST("/login", handler.Login)
	//public.GET("/validate", handler.Validate)
	v1 := router.Group("/api/v1")
	v1.Use(middleware.RequireAuth)
	{
		v1.GET("/validate", handler.Validate)
		v1.POST("user", handlers.CreateUser)
		v1.GET("user", handlers.GetAllUsers)
		v1.GET("user/:id", handlers.GetOneUser)
		v1.DELETE("user/:id", handlers.DeleteUser)
		v1.PUT("user/:id", handlers.UpdateUser)

	}

	return router
}
