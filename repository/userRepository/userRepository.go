package userRepository

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (us *UserRepository) GetUserDetailWithId(id user.Id) (*schema.User, error) {
	var user schema.User
	tx := us.db.Where("uuid = ?", id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (us *UserRepository) GetUserPreferences(id user.Id) (*schema.UserPreferences, error) {
	var userPreferences schema.UserPreferences
	tx := us.db.Where("user_uuid = ?", id).First(&userPreferences)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &userPreferences, nil
}

func (us *UserRepository) UpdateUserDetail(ctx echo.Context, username string, id user.Id) (*schema.User, error) {
	var user schema.User
	tx := us.db.Model(&user).Where("uuid= ?", id).Update("username", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (us *UserRepository) UpdateUserPreferences(ctx echo.Context, body *user.UpdateUserPreferencesJSONBody, id user.Id) (*schema.UserPreferences, error) {
	var userPreferences schema.UserPreferences
	tx := us.db.Model(&userPreferences).Where("user_uuid = ?", id).Update("preferences", body.Preference)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &userPreferences, nil
}
