package handlers

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// It accepts a *gin.Context object as a parameter for handling the HTTP request and response.
// This endpoint is used to get a list of all users.
// @Summary Get all users
// @Description Retrieves a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} Model.User
// @Router /api/v1/users [get]
func GetAllUsers(req *gin.Context) {
	var tasks []Model.User

	err := services.GetAllUsers(&tasks)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, tasks)
	}
}
