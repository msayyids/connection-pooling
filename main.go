package main

import (
	"pgxpool/config"
	"pgxpool/controller"
	"pgxpool/repo"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDb()
	repo := repo.Repo{DB: db}
	controller := controller.Controller{Ctrl: repo}

	e := echo.New()

	e.POST("/users", controller.Register)

	e.Logger.Fatal(e.Start(":8080"))

}
