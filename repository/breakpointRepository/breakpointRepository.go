package breakpointRepository

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewBreakpointRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
