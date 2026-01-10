package schema

import "time"

type BreakPointGenerateHistory struct {
	Id         int64     `gorm:"primaryKey;"`
	UserID     int64     `gorm:"not null;unique"`
	Created_at time.Time `gorm:"not null"`
}
