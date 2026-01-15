package plansUsecase

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/plans"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/KaungHtetMon29/BreakPoint_Backend/repository"
	"github.com/labstack/echo/v4"
)

type PlansUsecase struct {
	plansRepo repository.PlansRepository
}

func NewPlansUsecase(plansRepo repository.PlansRepository) *PlansUsecase {
	return &PlansUsecase{
		plansRepo: plansRepo,
	}
}

func (pu *PlansUsecase) GetCurrentPlan(ctx echo.Context, id plans.Id) (*schema.UserPlans, error) {
	plan, err := pu.plansRepo.GetCurrentPlan(ctx, id)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (pu *PlansUsecase) GetPlanHistory(ctx echo.Context, id plans.Id) ([]schema.UserPlans, error) {
	plans, err := pu.plansRepo.GetPlanHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	return plans, nil
}
