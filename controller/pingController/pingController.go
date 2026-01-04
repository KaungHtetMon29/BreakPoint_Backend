package pingController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Ping struct {
}

func NewPingCtrler() *Ping {
	return &Ping{}
}

func (pc *Ping) Getping(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "fine ping")
}
