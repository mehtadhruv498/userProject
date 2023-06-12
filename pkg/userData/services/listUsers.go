package services

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/store"
)

func GetAllUsers(users *[]Model.User) (err error) {

	err = store.GetAllUsers(users)
	if err != nil {
		return err
	}
	return nil
}
