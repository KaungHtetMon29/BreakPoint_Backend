package main

import (
	"time"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/public/ping"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/public/ping1"
	p1controller "github.com/KaungHtetMon29/BreakPoint_Backend/controller/ping1Controller"
	pcontroller "github.com/KaungHtetMon29/BreakPoint_Backend/controller/pingController"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	e := echo.New()
	dsn := "host=localhost user=test password=testkhm dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(
		&schema.Admin{},
		&schema.User{},
		&schema.BreakPointGenerateHistory{},
		&schema.UserPreferences{},
		&schema.UserPlanHistory{},
		&schema.BreakPointTechniques{},
	)
	psql, err := db.DB()
	if err != nil {
		panic("cannot get database object")
	}
	psql.SetMaxIdleConns(10)
	psql.SetMaxOpenConns(100)
	psql.SetConnMaxLifetime(30 * time.Minute)
	p1Ctrler := p1controller.NewPing1Ctrler()
	pCtrler := pcontroller.NewPingCtrler()
	// admin := e.Group("/admin")
	user := e.Group("/public")
	ping1.RegisterHandlers(user, p1Ctrler)
	ping.RegisterHandlers(user, pCtrler)
	e.Logger.Fatal(e.Start(":1323"))
}
