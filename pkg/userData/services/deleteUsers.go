package services

import (
	config "example/userProject/Config"
	Model "example/userProject/pkg/userData/Models"
)

func DeleteUser(user *Model.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)

	return nil
}
