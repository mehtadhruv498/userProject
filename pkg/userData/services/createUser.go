package services

import (
	Model "example/userproject/pkg/userData/Models"
	"example/userproject/pkg/userData/store"
)

func CreateUser(user *Model.User) (err error) {
	err = store.CreateUser(user)
	if err != nil {
		return err
	}
	return nil

}
