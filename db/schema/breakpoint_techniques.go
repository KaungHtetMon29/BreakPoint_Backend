package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type BreakPointTechniques struct {
	Id         int64     `gorm:"primaryKey;"`
	UUID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique"`
	UserUUID   uuid.UUID `gorm:"not null;unique"`
	Is_active  bool      `gorm:"not null;default:false"`
	Created_at time.Time `gorm:"not null"`
	Deleted_at time.Time
	Technique  datatypes.JSON
}
