package Model

import "gorm.io/gorm"

type UserInput struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
