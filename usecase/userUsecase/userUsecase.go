package userUsecase

import (
	"fmt"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (us *UserUsecase) GetUserDetail(id user.Id) (*schema.User, error) {
	fmt.Println("USECASE CALLED")
	user, err := us.userRepo.GetUserDetailWithId(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserUsecase) GetUserPreferences(id user.Id) (*schema.UserPreferences, error) {
	userPreferences, err := us.userRepo.GetUserPreferences(id)
	if err != nil {
		return nil, err
	}
	return userPreferences, nil
}
