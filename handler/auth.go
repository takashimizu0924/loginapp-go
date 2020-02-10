package handler

import (
	"net/http"

	"../model"
	"github.com/labstack/echo"
)

// func Login(c echo.Context) error {

// }

func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Name == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "invalid name or password",
		}
	}
	if u := model.FindUser(&model.User{Name: user.Name}); u.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exist",
		}
	}

	model.CreateUser(user)
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}
