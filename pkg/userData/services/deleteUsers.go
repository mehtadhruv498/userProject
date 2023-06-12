package services

import (
	Model "example/userProject/pkg/userData/Models"
	"example/userProject/pkg/userData/store"
)

func DeleteUser(user *Model.User, id string) (err error) {
	err = store.GetOneUser(user,id)

	if err != nil {
		return err
	}
	err = store.DeleteUser(user, id)
	if err != nil {
		return err
	}
	return nil
}
