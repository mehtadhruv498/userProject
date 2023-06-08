package handlers

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
