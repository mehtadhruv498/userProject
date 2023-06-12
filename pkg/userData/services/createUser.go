package services

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/store"
)

func CreateUser(user *Model.User) (err error) {
	err = store.CreateUser(user)
	if err != nil {
		return err
	}
	return nil

}
