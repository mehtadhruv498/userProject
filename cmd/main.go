package main

import (
	config "example/userProject/Config"
	"fmt"

	//"github.com/gin-gonic/gin"
	routes "example/userProject/pkg/userData"
	Model "example/userProject/pkg/userData/Models"
)

func init() {
	//config.LoadEnvVariables()
	config.ConnectToDB()
}

// @title Go + Gin + Gorm User API
// @version 1.0
// @description This is a sample Rest API server.
// @host localhost:3000
// @BasePath /
func main() {
	fmt.Println("This is a Rest Api Project")
	config.DB.AutoMigrate(&Model.User{})
	router := routes.SetupRouter()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run(":3000")

}
