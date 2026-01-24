package breakpointUsecase

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/breakpoints"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository"
	"github.com/labstack/echo/v4"
)

type BreakpointUsecase struct {
	breakpointRepo repository.BreakpointRepository
}

func NewBreakpointUsecase(breakpointRepo repository.BreakpointRepository) *BreakpointUsecase {
	return &BreakpointUsecase{
		breakpointRepo: breakpointRepo,
	}
}

func (bp *BreakpointUsecase) GenerateBreakPoint(ctx echo.Context) error {
	err := bp.breakpointRepo.GenerateBreakPoint(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (bp *BreakpointUsecase) GetBreakPointTechniques(ctx echo.Context, id breakpoints.Id) ([]schema.BreakPointTechniques, error) {
	techiques, err := bp.breakpointRepo.GetBreakPointTechniques(ctx, id)
	if err != nil {
		return nil, err
	}
	return techiques, nil
}

func (bp *BreakpointUsecase) GetBreakPointHistory(ctx echo.Context, id breakpoints.Id) ([]schema.BreakPointGenerateHistory, error) {
	bpGenHistory, err := bp.breakpointRepo.GetBreakPointHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	return bpGenHistory, nil
}
