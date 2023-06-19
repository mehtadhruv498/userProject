package routes

import (
	handler "example/userproject/pkg/userAuth/handlers"
	"example/userproject/pkg/userAuth/middleware"

	//"example/userproject/pkg/userAuth/middleware"
	"example/userproject/pkg/userData/handlers"
	"time"

	//"time"

	//"github.com/gin-contrib/cors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
	//"go.elastic.co/apm/module/apmgin"
)

// Routes function to serve endpoints
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(apmgin.Middleware(router))
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-BrowserFingerprint", "X-Workspace-ID"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	public := router.Group("/api")
	public.POST("/register", handler.Signup)
	public.POST("/login", handler.Login)
	v1 := router.Group("/api/v1")
	v1.Use(middleware.RequireAuth)
	{
		
		v1.POST("user", handlers.CreateUser)
		v1.GET("user", handlers.GetAllUsers)
		v1.GET("user/:id", handlers.GetOneUser)
		v1.DELETE("user/:id", handlers.DeleteUser)
		v1.PUT("user/:id", handlers.UpdateUser)

	}

	return router
}
