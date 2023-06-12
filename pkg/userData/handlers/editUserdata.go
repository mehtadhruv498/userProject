package handlers

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateUser godoc
// @Summary      Update User details
// @Description  Update the details of a user with a particular id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param id path string true "User ID"
// @Success      200 {string} string "Updated User Successfully"
// @Router       /api/v1/users/:id [put]
func UpdateUser(req *gin.Context) {
	var user Model.User

	id := req.Params.ByName("id")

	req.BindJSON(&user)
	err := services.UpdateUser(&user, id)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, user)
	}
}
