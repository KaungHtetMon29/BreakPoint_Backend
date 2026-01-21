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

type Detail struct {
	AlarmTime string `json:"alarm_time" jsonschema:"title=alarm_time,example= 9 am" jsonschema_description:"timer start time for start time. no additional messages or comments. strictly follow the example format"`
	Label     string `json:"label" jsonschema:"title=label,example= take a break/start work" jsonschema_description:"label for that alarm"`
}

type AlarmPattern struct {
	StartTime Detail `json:"start_time" jsonschema:"title=start_time,example= 9 am" jsonschema_description:"timer start time for start time. no additional messages or comments. strictly follow the example format"`
	StopTime  Detail `json:"stop_time" jsonschema:"title=stop_time,example= 9:30 am" jsonschema_description:"timer start time for stop time. no additional messages or comments. strictly follow the example format"`
	//can add break reason(walk, snack etc)
}

type AlarmSchedule struct {
	AlarmPatterns []AlarmPattern `json:"alarm_patterns" jsonschema:"title=alarm_patterns,type=array"`
}

type VariousSchedules struct {
	Variations []AlarmSchedule `json:"variations" jsonschema:"title=variations,type=array"`
}

// type TestSchema struct {
// 	WorkDuration  int `json:"work_duration" jsonschema:"title=work_duration,example=3600" jsonschema_description:"the working duration in seconds"`
// 	BreakDuration int `json:"break_duration" jsonschema:"title=break_duration, example=300" jsonschema_description:"the break duration in seconds"`
// }

// type TestSchemaArr struct {
// 	BreakTechniques []TestSchema `json:"break_techniques" jsonschema:"title=break_techniques,type=array"`
// }

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

var alarmSchedule = GenerateSchema[VariousSchedules]()

func main() {
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        "work_and_break_time_suggestion",
		Description: openai.String("Work and break time suggestion"),
		Schema:      alarmSchedule,
		Strict:      openai.Bool(true),
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_KEY")),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: schemaParam,
			},
		},
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.AssistantMessage(`
				Based on health analysis and research, suggest a work and break schedule for a software engineer with ADHD. 
				Details:
				- Age: 24
				- Work hours: 8 hours (7:30 am - 4:30 pm)
				- Lunch break: 1 hour
				- Health conditions: eye strain, headache, back pain
				- Goal: Increase productivity and reduce stress
				- Suggestion: Provide Various Break Techniques based on the Health conditions, the content you know and create your own break schedule version think out of the box.
				- Settings for break time generation: break time can have range from 5 mins to 20 mins.
				- variations: 3 schedules types various break schedules
				Provide the response in JSON format with the following structure:
				[
				{"start_time": {""alarm_time":9:00, "label":"start work"},"stop_time": {""alarm_time":9:30, "label":"take a break"}},
				{"start_time": {""alarm_time":9:00, "label":"start work"},"stop_time": {""alarm_time":9:30, "label":"take a break"}},
				{"start_time": {""alarm_time":9:00, "label":"start work"},"stop_time": {""alarm_time":9:30, "label":"take a break"}}, and so on
				]
				note: use only minute unit for the times. don't miss lunch break time.
				Time should be in 24 hour format
				Do not include any additional comments or messages.
			`),
		},
		Model: openai.ChatModelGPT5ChatLatest,
		N:     openai.Int(1),
	})
	if err != nil {
		panic(err.Error())
	}
	var catorigin AlarmSchedule
	_ = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &catorigin)
	println(chatCompletion.Choices[0].Message.Content)
	// println(chatCompletion.Choices[1].Message.Content)
	// println(chatCompletion.Choices[2].Message.Content)
	// println(catorigin.BreakTechniques[0].BreakDuration)
	// println(catorigin.BreakTechniques[1].WorkDuration)
	// println(catorigin.BreakTechniques[1].BreakDuration)
	// println(catorigin.BreakTechniques[2].WorkDuration)
	// println(catorigin.BreakTechniques[2].BreakDuration)
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
