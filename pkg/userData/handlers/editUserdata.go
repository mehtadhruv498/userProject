package handlers

import (
	Model "example/userproject/pkg/userData/Models"
	"example/userproject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateUser godoc
// @Summary      Update User details
// @Description  Update the details of a user with a particular id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body Model.User true "User"
// @Success      200 {string} string "Updated User Successfully"
// @Router       /api/v1/user/{id} [put]
func UpdateUser(req *gin.Context) {
	var user Model.User

	id := req.Params.ByName("id")

	req.BindJSON(&user)
	err := services.UpdateUser(&user, id)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
		return
	}
	req.JSON(http.StatusOK, user)
}
