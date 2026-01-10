package schema

import (
	"time"
)

type Plan string

const (
	Free    Plan = "free"
	Premium Plan = "premium"
)

type UserPlanHistory struct {
	Id         int64 `gorm:"primaryKey;"`
	UserID     int64 `gorm:"not null;unique"`
	PlanType   Plan  `gorm:"not null"`
	Start_Date time.Time
	End_Date   time.Time
	Created_at time.Time `gorm:"not null"`
	Deleted_at time.Time
}
