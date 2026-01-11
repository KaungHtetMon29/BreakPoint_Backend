package schema

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	Id         int64     `gorm:"primaryKey;autoIncrement;"`
	UUID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique"`
	Username   string    `gorm:"not null"`
	Email      string    `gorm:"unique"`
	Level      string    `gorm:"not null"`
	Created_at time.Time `gorm:"not null"`
	Deleted_at time.Time
}
