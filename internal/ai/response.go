package ai

import (
	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go/v3"
)

type AlarmDetail struct {
	AlarmTime string `json:"alarm_time" jsonschema:"title=alarm_time,example= 9 am" jsonschema_description:"timer start time for start time. no additional messages or comments. strictly follow the example format"`
	Label     string `json:"label" jsonschema:"title=label,example= take a break/start work" jsonschema_description:"label for that alarm"`
}

type AlarmPattern struct {
	StartTime AlarmDetail `json:"start_time" jsonschema:"title=start_time,example= 9 am" jsonschema_description:"timer start time for start time. no additional messages or comments. strictly follow the example format"`
	StopTime  AlarmDetail `json:"stop_time" jsonschema:"title=stop_time,example= 9:30 am" jsonschema_description:"timer start time for stop time. no additional messages or comments. strictly follow the example format"`
	//can add break reason(walk, snack etc)
}

type AlarmSchedule struct {
	AlarmPatterns []AlarmPattern `json:"alarm_patterns" jsonschema:"title=alarm_patterns,type=array"`
}

type VariousSchedules struct {
	Variations []AlarmSchedule `json:"variations" jsonschema:"title=variations,type=array"`
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

var alarmSchedule = GenerateSchema[VariousSchedules]()

var SchemaParam = openai.ResponseFormatJSONSchemaJSONSchemaParam{
	Name:        "work_and_break_time_suggestion",
	Description: openai.String("Work and break time suggestion"),
	Schema:      alarmSchedule,
	Strict:      openai.Bool(true),
}
