package userRepository

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
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
