package services

import (
	config "example/userProject/Config"
	Model "example/userProject/pkg/userData/Models"
)

func GetAllUsers(user *[]Model.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}

	return nil
}
