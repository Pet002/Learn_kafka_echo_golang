package handler

import (
	"time"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	User User
}

func NewUserHandler(user User) *userHandler {
	return &userHandler{
		User: user,
	}
}

func (u *userHandler) GetHandler(c echo.Context) error {
	time.Sleep(3 * time.Second)
	return c.JSONPretty(200, u.User, "")
}
