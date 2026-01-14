package repository

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/breakpoints"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	GetUserDetailWithId(id user.Id) (*schema.User, error)
	GetUserPreferences(id user.Id) (*schema.UserPreferences, error)
}

type BreakpointRepository interface {
	GetBreakPointTechniques(ctx echo.Context, id breakpoints.Id) ([]schema.BreakPointTechniques, error)
	GetBreakPointHistory(ctx echo.Context, id breakpoints.Id) ([]schema.BreakPointGenerateHistory, error)
}
