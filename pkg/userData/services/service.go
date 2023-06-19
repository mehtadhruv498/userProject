package services

import (
	Model "example/userproject/pkg/userData/Models"
	"example/userproject/pkg/userData/store"
)

type UserService interface {
	CreateUser(user *Model.User) error
	UpdateUser(user *Model.User, id string) error
	GetAllUsers(users *[]Model.User) error
	GetOneUser(user *Model.User, id string) (error, Model.User)
	DeleteUser(user *Model.User, id string) error
}

type Service struct {
	store store.Store
}

func NewUserService(store store.Store) *Service {
	return &Service{
		store: store,
	}
}
