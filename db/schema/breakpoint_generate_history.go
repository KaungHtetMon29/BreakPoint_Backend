package schema

import (
	"time"

	"github.com/google/uuid"
)

type BreakPointGenerateHistory struct {
	Id         int64     `gorm:"primaryKey;"`
	UUID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique"`
	UserID     int64     `gorm:"not null;unique"`
	Created_at time.Time `gorm:"not null"`
}
