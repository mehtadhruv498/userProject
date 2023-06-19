package handlers

import (
	Model "example/userproject/pkg/userData/Models"
	"example/userproject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOneUser godoc
// It accepts a *gin.Context object as a parameter for handling the HTTP request and response.
// This endpoint is used to get details of a particular user.
// @Summary Get a particluar user
// @Description Gets details of a particular user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} Model.User
// @Router /api/v1/user/{id} [get]
func GetOneUser(req *gin.Context) {
	var user Model.User

	id := req.Params.ByName("id")

	err := services.GetOneUser(&user, id)

	if err != nil {
		req.AbortWithStatus(http.StatusNotFound)
	} else {
		req.JSON(http.StatusOK, user)
	}
}
