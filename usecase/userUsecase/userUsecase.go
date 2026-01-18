package userUsecase

import (
	"fmt"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository"
	"github.com/labstack/echo/v4"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (us *UserUsecase) GetUserDetail(ctx echo.Context, id user.Id) (*schema.User, error) {
	fmt.Println("USECASE CALLED")
	user, err := us.userRepo.GetUserDetailWithId(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserUsecase) GetUserPreferences(ctx echo.Context, id user.Id) (*schema.UserPreferences, error) {
	userPreferences, err := us.userRepo.GetUserPreferences(id)
	if err != nil {
		return nil, err
	}
	return userPreferences, nil
}

func (us *UserUsecase) UpdateUserDetail(ctx echo.Context, body *user.UpdateUserDetailJSONRequestBody, id user.Id) (*schema.User, error) {
	// add validation
	user, err := us.userRepo.UpdateUserDetail(ctx, *body.Username, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserUsecase) UpdateUserPreferences(ctx echo.Context, body *user.UpdateUserPreferencesJSONBody, id user.Id) (*schema.UserPreferences, error) {
	// add validation
	userPreferences, err := us.userRepo.UpdateUserPreferences(ctx, body, id)
	if err != nil {
		return nil, err
	}
	return userPreferences, nil
}
