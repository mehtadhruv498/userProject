package handlers

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(req *gin.Context) {
	var tasks []Model.User

	err := services.GetAllUsers(&tasks)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, tasks)
	}
}
