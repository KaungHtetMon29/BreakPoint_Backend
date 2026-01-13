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

type UserPlanHistory struct {
	Id         int64     `gorm:"primaryKey;"`
	UUID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique"`
	UserUUID   int64     `gorm:"not null;unique"`
	PlanType   Plan      `gorm:"not null"`
	Start_Date time.Time
	End_Date   time.Time
	Created_at time.Time `gorm:"not null"`
	Deleted_at time.Time
}
