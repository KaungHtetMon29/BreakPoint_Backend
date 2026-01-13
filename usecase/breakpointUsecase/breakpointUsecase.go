package breakpointUsecase

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository"
)

type BreakpointUsecase struct {
	breakpointRepo repository.BreakpointRepository
}

func NewBreakpointUsecase(userRepo repository.UserRepository) *BreakpointUsecase {
	return &BreakpointUsecase{}
}

// func (bp *BreakpointUsecase) GetBreakPointTechniques(ctx echo.Context) error {}
