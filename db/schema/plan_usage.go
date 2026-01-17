package schema

import (
	"time"

	"github.com/google/uuid"
)

type PlanUsage struct {
	Id            int64     `gorm:"primaryKey;"`
	UUID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique"`
	UserUUID      uuid.UUID `gorm:"type:uuid;not null"`
	PlanUUID      uuid.UUID `gorm:"not null;"`
	GenerateCount int64     `gorm:"not null;default: 0"`
	Created_at    time.Time `gorm:"not null"`
}
