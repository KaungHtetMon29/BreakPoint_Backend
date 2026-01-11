package userController

import (
	"net/http"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/labstack/echo/v4"
)

type User struct {
}

func NewUserCtrler() *User {
	return &User{}
}

func (pc *User) GetUserDetail(ctx echo.Context, id user.Id) error {
	return ctx.JSON(http.StatusOK, "get user detail")
}

func (pc *User) UpdateUserDetail(ctx echo.Context, id user.Id) error {
	return ctx.JSON(http.StatusOK, "update user detail")
}

func (pc *User) GetUserPreferences(ctx echo.Context, id user.Id) error {
	return ctx.JSON(http.StatusOK, "get user preference")
}

func (pc *User) UpdateUserPreferences(ctx echo.Context, id user.Id) error {
	return ctx.JSON(http.StatusOK, "update user preference")
}
