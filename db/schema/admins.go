package schema

import "time"

type Admin struct {
	Id         int64     `gorm:"primaryKey;autoIncrement;"`
	Username   string    `gorm:"not null"`
	Email      string    `gorm:"unique"`
	Level      string    `gorm:"not null"`
	Created_at time.Time `gorm:"not null"`
	Deleted_at time.Time
}
