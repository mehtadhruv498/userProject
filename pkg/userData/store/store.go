package store

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

func DeleteUser(user *Model.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)

	return nil
}

func UpdateUser(user *Model.User) (err error) {
	config.DB.Save(user)
	return nil
}

func GetOneUser(user *Model.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers(user *[]Model.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}

	return nil
}
