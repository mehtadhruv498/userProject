package handlers

import "example/userproject/pkg/userData/services"

type Handler struct {
	Service services.UserService
}

var userService services.UserService

func SetUserService(user services.UserService) {
	userService = user
}

func New() *Handler {
	return &Handler{
		Service: userService,
	}
}
