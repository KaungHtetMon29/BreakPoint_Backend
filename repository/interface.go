package repository

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
)

type UserRepository interface {
	GetUserDetailWithId(id user.Id) (*schema.User, error)
	GetUserPreferences(id user.Id) (*schema.UserPreferences, error)
}

type BreakpointRepository interface {
}
