package ping1Controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Ping1 struct {
}

func NewPing1Ctrler() *Ping1 {
	return &Ping1{}
}

func (pc *Ping1) Ping1(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "fine")
}
