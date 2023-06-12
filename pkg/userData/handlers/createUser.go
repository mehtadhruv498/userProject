package handlers

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary      Create a User
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body Model.User true "User"
// @Success      200 {string} string "Created User Successfully"
// @Router       /api/v1/users [post]
func CreateUser(req *gin.Context) {
	var user Model.User

	if err := req.ShouldBindJSON(&user); err != nil {
		req.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateUser(&user)
	if err != nil {
		req.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	req.JSON(http.StatusOK, gin.H{"user": "Created User Successfully"})
}
