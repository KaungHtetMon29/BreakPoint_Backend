package breakpointsController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Breakpoints struct {
}

func NewBreakpointsCtrler() *Breakpoints {
	return &Breakpoints{}
}

func (pc *Breakpoints) GenerateBreakPoint(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Generate Breakpoint")
}

func (pc *Breakpoints) GetBreakPointHistory(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Get Break Point History")
}

func (pc *Breakpoints) GetBreakPointTechniques(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Get Break Point Techniques")
}
