package main

import (
	"time"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/admin"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/auth"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/breakpoints"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/plans"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/public/ping"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/public/ping1"
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/user"
	"github.com/KaungHtetMon29/BreakPoint_Backend/controller/adminController"
	"github.com/KaungHtetMon29/BreakPoint_Backend/controller/authController"
	"github.com/KaungHtetMon29/BreakPoint_Backend/controller/breakpointsController"
	p1controller "github.com/KaungHtetMon29/BreakPoint_Backend/controller/ping1Controller"
	pcontroller "github.com/KaungHtetMon29/BreakPoint_Backend/controller/pingController"
	"github.com/KaungHtetMon29/BreakPoint_Backend/controller/plansController"
	"github.com/KaungHtetMon29/BreakPoint_Backend/controller/userController"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository/breakpointRepository"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository/plansRepository"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository/userRepository"
	"github.com/KaungHtetMon29/BreakPoint_Backend/usecase/breakpointUsecase"
	plansUsecase "github.com/KaungHtetMon29/BreakPoint_Backend/usecase/plans"
	"github.com/KaungHtetMon29/BreakPoint_Backend/usecase/userUsecase"
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
	err = db.AutoMigrate(
		&schema.Admin{},
		&schema.User{},
		&schema.BreakPointGenerateHistory{},
		&schema.UserPreferences{},
		&schema.UserPlans{},
		&schema.BreakPointTechniques{},
		&schema.PlanUsage{},
	)
	if err != nil {
		panic("automigration failed")
	}
	psql, err := db.DB()
	if err != nil {
		panic("cannot get database object")
	}
	psql.SetMaxIdleConns(10)
	psql.SetMaxOpenConns(100)
	psql.SetConnMaxLifetime(30 * time.Minute)
	adminGroup := e.Group("/admin")
	userGroup := e.Group("/user")
	authGroup := e.Group("/auth")
	breakPointGroup := e.Group("/breakpoints")
	planGroup := e.Group("/plans")

	userRepo := userRepository.NewUserRepository(db)
	breakpointRepo := breakpointRepository.NewBreakpointRepository(db)
	plansRepo := plansRepository.NewPlansRepository(db)

	userUsecase := userUsecase.NewUserUsecase(userRepo)
	breakpointUsecase := breakpointUsecase.NewBreakpointUsecase(breakpointRepo)
	plansUsecase := plansUsecase.NewPlansUsecase(plansRepo)

	p1Ctrler := p1controller.NewPing1Ctrler()
	pCtrler := pcontroller.NewPingCtrler()
	userCtrl := userController.NewUserCtrler(userUsecase)
	authCtrl := authController.NewAuthCtrler()
	bpCtrl := breakpointsController.NewBreakpointsCtrler(breakpointUsecase)
	planCtrl := plansController.NewPlansCtrler(plansUsecase)
	adminCtrl := adminController.NewAdminCtrler()

	admin.RegisterHandlers(adminGroup, adminCtrl)
	ping1.RegisterHandlers(userGroup, p1Ctrler)
	ping.RegisterHandlers(userGroup, pCtrler)
	user.RegisterHandlers(userGroup, userCtrl)
	auth.RegisterHandlers(authGroup, authCtrl)
	plans.RegisterHandlers(planGroup, planCtrl)
	breakpoints.RegisterHandlers(breakPointGroup, bpCtrl)

	e.Logger.Fatal(e.Start(":1323"))
}
