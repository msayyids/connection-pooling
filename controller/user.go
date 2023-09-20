package controller

import (
	"log"
	"net/http"
	"pgxpool/helper"
	"pgxpool/model"
	"pgxpool/repo"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Ctrl repo.Repo
}

func (cn *Controller) Register(c echo.Context) error {

	var registerRequest model.User
	if err := c.Bind(&registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	newUser := model.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: helper.HashedPassword(registerRequest.Password),
	}

	registeredUser, err := cn.Ctrl.AddUser(newUser)
	if err != nil {

		log.Println("Error during user registration:", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User registered successfully",
		"Name":    registeredUser.Name,
	})
}
