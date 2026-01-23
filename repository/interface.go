package repository

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/breakpoints"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/plans"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	GetUserDetailWithId(id user.Id) (*schema.User, error)
	GetUserPreferences(id user.Id) (*schema.UserPreferences, error)
	UpdateUserDetail(ctx echo.Context, username string, id user.Id) (*schema.User, error)
	UpdateUserPreferences(ctx echo.Context, body *user.UpdateUserPreferencesJSONBody, id user.Id) (*schema.UserPreferences, error)
}

type BreakpointRepository interface {
	GenerateBreakPoint(ctx echo.Context) error
	GetBreakPointTechniques(ctx echo.Context, id breakpoints.Id) ([]schema.BreakPointTechniques, error)
	GetBreakPointHistory(ctx echo.Context, id breakpoints.Id) ([]schema.BreakPointGenerateHistory, error)
}

type PlansRepository interface {
	GetCurrentPlan(ctx echo.Context, id plans.Id) (*schema.UserPlans, error)
	GetPlanHistory(ctx echo.Context, id plans.Id) ([]schema.UserPlans, error)
	GetPlanUsage(ctx echo.Context, id plans.Id) ([]schema.PlanUsage, error)
}
