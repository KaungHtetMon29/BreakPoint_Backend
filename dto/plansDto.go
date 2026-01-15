package dto

type PlanDto struct {
	UUID      string `json:"uuid"`
	PlanType  string `json:"plan_type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
