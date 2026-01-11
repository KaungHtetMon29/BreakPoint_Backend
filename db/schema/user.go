package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id                        int64     `gorm:"primaryKey;autoIncrement;"`
	UUID                      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique"`
	Username                  string    `gorm:"not null"`
	Email                     string    `gorm:"unique"`
	Created_at                time.Time `gorm:"not null"`
	Updated_at                time.Time
	Deleted_at                time.Time
	BreakPointGenerateHistory []BreakPointGenerateHistory `gorm:"foreignKey:UserID"`
	UserPreferences           []UserPreferences           `gorm:"foreignKey:UserID"`
	UserPlanHistory           []UserPlanHistory           `gorm:"foreignKey:UserID"`
}
