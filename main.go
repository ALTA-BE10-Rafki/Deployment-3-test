package main

import (
	"apigorm/controller/user"
	"apigorm/database/mysql"
	"apigorm/model"

	"github.com/labstack/echo/v4"
)

func main() {

	db := mysql.InitDB()
	mysql.MigrateData(db)
	e := echo.New()
	userModel := model.UserModel{DB: db}
	userController := user.UserController{Model: userModel}

	e.GET("/users", userController.GetAll())
	e.GET("/users/:id", userController.GetSpesificUser())
	e.POST("/users", userController.InsertUser())
	e.DELETE("/users/:id", userController.DeleteData())
	e.PUT("/users/:id", userController.UpdateUser())

	e.Logger.Fatal(e.Start(":8000"))
}
