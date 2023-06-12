package services

import (
	"example/userProject/pkg/userData/store"

	Model "example/userProject/pkg/userData/Models"
)

func GetOneUser(user *Model.User, id string) (err error) {

	err = store.GetOneUser(user, id)
	if err != nil {
		return err
	}
	return nil
}
