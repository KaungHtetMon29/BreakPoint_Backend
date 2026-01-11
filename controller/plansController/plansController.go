package plansController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Plans struct {
}

func NewPlansCtrler() *Plans {
	return &Plans{}
}

func (pc *Plans) GetCurrentPlan(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "get current plan")
}

func (pc *Plans) GetPlanHistory(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "get plan history")
}

func (pc *Plans) PostUpgradePlan(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "post upgrade plan")
}

func (pc *Plans) GetPlanUsage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "get plan usage")
}
