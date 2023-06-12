package services

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/store"
)

func UpdateUser(user *Model.User, id string) (err error) {
	err = store.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
