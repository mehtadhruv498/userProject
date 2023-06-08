package handlers

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(req *gin.Context) {
	var user Model.User

	// Bind the request body to the user struct
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
