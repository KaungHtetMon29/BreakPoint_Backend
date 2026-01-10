package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id                        int64     `gorm:"primaryKey;autoIncrement;"`
	Username                  string    `gorm:"not null"`
	Email                     string    `gorm:"unique"`
	Created_at                time.Time `gorm:"not null"`
	Updated_at                time.Time `gorm:"not null"`
	Deleted_at                time.Time
	BreakPointGenerateHistory []BreakPointGenerateHistory `gorm:"foreignKey:UserID"`
	UserPreferences           []UserPreferences           `gorm:"foreignKey:UserID"`
	UserPlanHistory           []UserPlanHistory           `gorm:"foreignKey:UserID"`
}
