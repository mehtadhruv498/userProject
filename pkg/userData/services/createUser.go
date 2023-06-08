package services

import (
	config "example/userProject/Config"
	Model "example/userProject/pkg/userData/Models"
)

func CreateUser(user *Model.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

