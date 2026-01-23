package ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go/v3"
)

func Request(client openai.Client) {
	fmt.Println("run")
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: SchemaParam,
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
}
