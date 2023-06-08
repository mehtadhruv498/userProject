package handlers

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(req *gin.Context) {
	var task Model.User

	id := req.Params.ByName("id")

	err := services.GetOneUser(&task, id)

	if err != nil {
		req.JSON(http.StatusNotFound, task)
	}

	req.BindJSON(&task)
	err = services.UpdateUser(&task)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, task)
	}
}
