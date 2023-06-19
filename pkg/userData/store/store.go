package store

import (
	config "example/userproject/Config"
	Model "example/userproject/pkg/userData/Models"
)

type Store interface {
	CreateUser(user Model.User) error
	UpdateUser(initialUser, updatedUser Model.User) error
	GetAllUsers() (error, []Model.User)
	GetOneUser(id int) (error, Model.User)
	DeleteUser(id int) error
}

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
