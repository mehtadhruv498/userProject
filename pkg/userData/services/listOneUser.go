package services

import (
	config "example/userProject/Config"
	Model "example/userProject/pkg/userData/Models"
)

func GetOneUser(user *Model.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
