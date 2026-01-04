package main

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/public/ping"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/public/ping1"
	p1controller "github.com/KaungHtetMon29/BreakPoint_Backend/controller/ping1Controller"
	pcontroller "github.com/KaungHtetMon29/BreakPoint_Backend/controller/pingController"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	p1Ctrler := p1controller.NewPing1Ctrler()
	pCtrler := pcontroller.NewPingCtrler()
	// admin := e.Group("/admin")
	user := e.Group("/public")
	ping1.RegisterHandlers(user, p1Ctrler)
	ping.RegisterHandlers(user, pCtrler)
	e.Logger.Fatal(e.Start(":1323"))
}
