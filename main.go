package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {

	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "username:mehta123@tcp(127.0.0.1:3306)/userData")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Mysql DB is connected")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defer db.Close()

}
