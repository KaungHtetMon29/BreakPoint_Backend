package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
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
	"github.com/invopop/jsonschema"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestSchema struct {
	WorkDuration  int `json:"work_duration" jsonschema:"title=work_duration,example=3600" jsonschema_description:"the working duration in seconds"`
	BreakDuration int `json:"break_duration" jsonschema:"title=break_duration, example=300" jsonschema_description:"the break duration in seconds"`
}

type TestSchemaArr struct {
	BreakTechniques []TestSchema `json:"break_techniques" jsonschema:"title=break_techniques,type=array"`
}

func GenerateSchema[T any]() interface{} {
	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

var TestResSchema = GenerateSchema[TestSchemaArr]()

func main() {
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "work_and_break_time_suggestion",
		Description: openai.String("Work and break time suggestion"),
		Schema:      TestResSchema,
		Strict:      openai.Bool(true),
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_KEY")),
	)
	_, err = client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.AssistantMessage(`
				userid = 1234
				data = didn't took a break of 5 mins at 11:00
				date = 21/1/2026
			`),
		},
		Model: openai.ChatModelGPT5Mini,
		N:     openai.Int(1),
	})
	if err != nil {
		panic(err.Error())
	}

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: schemaParam,
			},
		},
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.AssistantMessage(`
				userid = 1234
				age = 24
				job = software engineer
				work hours = 8 hours
				work time = 7:30 am - 4:30 pm
				lunch break = 1 hour
				heal condition = "all good. only eye strain. I have headache. Also, have back pain"
				request = "suggest me duration to work continuously and duration to take a break.
				I have adhd.
				based on some health analysis and various break techniques and make your own break techniques based on your data and research.
				I want to be more productive and less stress from work"
				break_techniques = 3 techniques
			`),
		},
		Model: openai.ChatModelGPT5Mini,
		N:     openai.Int(1),
	})
	if err != nil {
		panic(err.Error())
	}
	var catorigin TestSchemaArr
	_ = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &catorigin)
	println(catorigin.BreakTechniques[0].WorkDuration)
	println(catorigin.BreakTechniques[0].BreakDuration)
	println(catorigin.BreakTechniques[1].WorkDuration)
	println(catorigin.BreakTechniques[1].BreakDuration)
	println(catorigin.BreakTechniques[2].WorkDuration)
	println(catorigin.BreakTechniques[2].BreakDuration)
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
