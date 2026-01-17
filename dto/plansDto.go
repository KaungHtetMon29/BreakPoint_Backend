package dto

type PlanDto struct {
	UUID      string `json:"uuid"`
	PlanType  string `json:"plan_type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type PlanUsageDto struct {
	GenCount int64  `json:"generation_count"`
	Date     string `json:"date"`
}
