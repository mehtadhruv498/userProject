package services

import (
	Model "example/userproject/pkg/userData/Models"
	"example/userproject/pkg/userData/store"
)

func UpdateUser(user *Model.User, id string) (err error) {
	err = store.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
