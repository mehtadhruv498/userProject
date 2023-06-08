package services

import (
	config "example/userProject/Config"
	Model "example/userProject/pkg/userData/Models"
)

func UpdateUser(user *Model.User) (err error) {
	config.DB.Save(user)
	return nil
}
