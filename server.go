package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
	authentication "github.com/KaungHtetMon29/BreakPoint_Backend/internal/auth"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository/breakpointRepository"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository/plansRepository"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository/userRepository"
	"github.com/KaungHtetMon29/BreakPoint_Backend/usecase/breakpointUsecase"
	plansUsecase "github.com/KaungHtetMon29/BreakPoint_Backend/usecase/plans"
	"github.com/KaungHtetMon29/BreakPoint_Backend/usecase/userUsecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_KEY")),
	)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"}, // Include the full URL
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}, // Add necessary headers
		AllowCredentials: true,
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusOK)
			}
			return next(c)
		}
	})
	dsn := "host=localhost user=test password=testkhm dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// oauth test
	scopes := strings.Split(os.Getenv("SCOPES"), ",")
	oauth := authentication.NewOauth(
		os.Getenv("OAUTH_REDIRECT_URL"),
		os.Getenv("GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		scopes,
	)

	fmt.Printf("visit the url for the auth dialog: %v", oauth.AuthCodeUrl)
	e.POST("/login", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"redirect_url": oauth.AuthCodeUrl,
		})
	})
	e.GET("/auth/callback", func(c echo.Context) error {
		code := c.Request().FormValue("code")
		fmt.Println(c.Request().FormValue("code"))
		userInfo, err := oauth.GetGoogleUserInfo(code)
		if err != nil {
			return err
		}
		token, err := authentication.CreateJWTToken(userInfo)
		if err != nil {
			return err
		}
		cookie := new(http.Cookie)
		cookie.Name = "test"
		cookie.Value = *token
		cookie.Path = "/"
		cookie.HttpOnly = true
		cookie.Secure = false
		c.SetCookie(cookie)
		return c.Redirect(http.StatusPermanentRedirect, "http://localhost:3000")
	})

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
		panic(err)
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
	breakpointRepo := breakpointRepository.NewBreakpointRepository(db, &client)
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
