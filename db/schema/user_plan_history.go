package schema

import (
	"time"

	"github.com/google/uuid"
)

type Plan string

const (
	Free    Plan = "free"
	Premium Plan = "premium"
)

type UserPlans struct {
	Id         int64     `gorm:"primaryKey;"`
	UUID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique"`
	UserUUID   uuid.UUID `gorm:"not null;"`
	PlanType   Plan      `gorm:"not null"`
	IsActive   bool      `gorm:"not null; default:false"`
	Start_Date *time.Time
	End_Date   *time.Time
	Created_at time.Time `gorm:"not null"`
	Deleted_at *time.Time
}
