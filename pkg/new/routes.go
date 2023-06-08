package routes

import (
	"github.com/gin-gonic/gin"

	Controller "example/userProject/pkg/new/Controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// v1.GET("", Controllers.Home)
		// v1.GET("about", Controllers.About)

		v1.POST("task/create", Controller.CreateTask)
		v1.GET("task", Controller.GetAllTasks)
		v1.GET("task/:id", Controller.GetATask)
		v1.PUT("task/:id", Controller.UpdateATask)
		v1.DELETE("task/:id", Controller.DeleteATask)
	}

	return router
}



