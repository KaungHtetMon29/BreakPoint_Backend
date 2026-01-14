package dto

type BreakPointTechniquesDto struct {
	UUID       string `json:"uuid"`
	Is_active  bool   `json:"is_active"`
	Techniques string `json:"techniques"`
}
