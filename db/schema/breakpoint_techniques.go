package schema

import (
	"time"

	"gorm.io/datatypes"
)

type BreakPointTechniques struct {
	Id         int64     `gorm:"primaryKey;"`
	UserID     int64     `gorm:"not null;unique"`
	Is_active  bool      `gorm:"not null;default:false"`
	Created_at time.Time `gorm:"not null"`
	Deleted_at time.Time
	Technique  datatypes.JSON
}
