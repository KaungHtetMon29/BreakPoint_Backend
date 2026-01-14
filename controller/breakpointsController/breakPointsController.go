package breakpointsController

import (
	"net/http"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/breakpoints"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/KaungHtetMon29/BreakPoint_Backend/dto"
	"github.com/KaungHtetMon29/BreakPoint_Backend/usecase"
	"github.com/labstack/echo/v4"
)

type Breakpoints struct {
	breakpointUsecase usecase.BreakpointUsecase
}

func NewBreakpointsCtrler(breakpointUsecase usecase.BreakpointUsecase) *Breakpoints {
	return &Breakpoints{
		breakpointUsecase: breakpointUsecase,
	}
}

func (bpc *Breakpoints) GenerateBreakPoint(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Generate Breakpoint")
}

func (bpc *Breakpoints) GetBreakPointHistory(ctx echo.Context, id breakpoints.Id) error {
	bpGenHistory, err := bpc.breakpointUsecase.GetBreakPointHistory(ctx, id)
	if err != nil {
		return err
	}
	dto := convertToBreakPointGenerateHistoryDTO(bpGenHistory)
	return ctx.JSON(http.StatusOK, dto)
}

func (bpc *Breakpoints) GetBreakPointTechniques(ctx echo.Context, id breakpoints.Id) error {
	bptechniques, err := bpc.breakpointUsecase.GetBreakPointTechniques(ctx, id)
	if err != nil {
		return err
	}
	dto := convertToBreakPointTechniquesDTO(bptechniques)
	return ctx.JSON(http.StatusOK, dto)
}

func convertToBreakPointTechniquesDTO(breakpointTechniques []schema.BreakPointTechniques) []dto.BreakPointTechniquesDto {
	var breakpointTechniquesDto = make([]dto.BreakPointTechniquesDto, len(breakpointTechniques))
	for i := range breakpointTechniques {
		breakpointTechniquesDto[i].UUID = breakpointTechniques[i].UUID.String()
		breakpointTechniquesDto[i].Is_active = breakpointTechniques[i].Is_active
		breakpointTechniquesDto[i].Techniques = string(breakpointTechniques[i].Technique)
	}
	return breakpointTechniquesDto
}

func convertToBreakPointGenerateHistoryDTO(breakpointGenerateHistory []schema.BreakPointGenerateHistory) []dto.BreakPointGenerateHistoryDto {
	var breakpointGenerateHistoryDto = make([]dto.BreakPointGenerateHistoryDto, len(breakpointGenerateHistory))
	for i := range breakpointGenerateHistory {
		breakpointGenerateHistoryDto[i].UUID = breakpointGenerateHistory[i].UUID.String()
		breakpointGenerateHistoryDto[i].Created_at = breakpointGenerateHistory[i].Created_at.String()
	}
	return breakpointGenerateHistoryDto
}
