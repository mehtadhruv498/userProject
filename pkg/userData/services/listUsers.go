package services

import (
	Model "example/userproject/pkg/userData/Models"
	"example/userproject/pkg/userData/store"
)

func GetAllUsers(users *[]Model.User) (err error) {

	err = store.GetAllUsers(users)
	if err != nil {
		return err
	}
	return nil
}
