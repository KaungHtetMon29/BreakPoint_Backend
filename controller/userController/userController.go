package userController

import (
	"net/http"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/dto"
	"github.com/KaungHtetMon29/BreakPoint_Backend/usecase"
	"github.com/labstack/echo/v4"
)

type User struct {
	userUsecase usecase.UserUsecase
}

func NewUserCtrler(userUsecase usecase.UserUsecase) *User {
	return &User{
		userUsecase: userUsecase,
	}
}

func (pc *User) GetUserDetail(ctx echo.Context, id user.Id) error {
	user, err := pc.userUsecase.GetUserDetail(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, dto.UserDto{
		UUID:     user.UUID.String(),
		UserName: user.Username,
		Email:    user.Email,
	})
}

func (pc *User) UpdateUserDetail(ctx echo.Context, id user.Id) error {
	return ctx.JSON(http.StatusOK, "update user detail")
}

func (pc *User) GetUserPreferences(ctx echo.Context, id user.Id) error {
	userPreferences, err := pc.userUsecase.GetUserPreferences(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, dto.UserPreferenceDto{
		Preference: string(userPreferences.Preferences),
	})
}

func (pc *User) UpdateUserPreferences(ctx echo.Context, id user.Id) error {
	return ctx.JSON(http.StatusOK, "update user preference")
}
