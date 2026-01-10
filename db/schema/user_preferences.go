package schema

import (
	"time"

	"gorm.io/datatypes"
)

type UserPreferences struct {
	Id          int64     `gorm:"primaryKey;"`
	UserID      int64     `gorm:"not null;unique"`
	Created_at  time.Time `gorm:"not null"`
	Deleted_at  time.Time
	Preferences datatypes.JSON
}
