package authController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
}

func NewAuthCtrler() *User {
	return &User{}
}

func (pc *User) Login(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "login")
}

func (pc *User) Logout(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "logout")
}

func (pc *User) SignUp(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "signup")
}

func (pc *User) GetProfile(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "get profile")
}
