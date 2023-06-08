package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "root:password@tcp(localhost:3306)/userData2?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database" + err.Error())
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get underlying *sql.DB")
	}
	DB.AutoMigrate()
	fmt.Println("DB connected at: ", sqlDB)
}
