package handlers

import (
	Model "example/userproject/pkg/userData/Models"
	"example/userproject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary      Delete a User
// @Description  Delete a user with the id given
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Success      200 {string} string "Deleted User Successfully"
// @Router       /api/v1/user/{id} [delete]
func DeleteUser(req *gin.Context) {
	var user Model.User

	id := req.Params.ByName("id")

	err := services.DeleteUser(&user, id)

	if err != nil {
		req.AbortWithStatus(http.StatusNotFound)
	} else {
		req.JSON(http.StatusOK, user)
	}
}
