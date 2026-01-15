package plansController

import (
	"net/http"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/plans"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/KaungHtetMon29/BreakPoint_Backend/dto"
	"github.com/KaungHtetMon29/BreakPoint_Backend/usecase"
	"github.com/labstack/echo/v4"
)

type Plans struct {
	plansUsecase usecase.PlansUsecase
}

func NewPlansCtrler(plansUsecase usecase.PlansUsecase) *Plans {
	return &Plans{
		plansUsecase: plansUsecase,
	}
}

func (pc *Plans) GetCurrentPlan(ctx echo.Context, id plans.Id) error {
	plan, err := pc.plansUsecase.GetCurrentPlan(ctx, id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, convertToCurrentPlanDto(plan))
}

func (pc *Plans) GetPlanHistory(ctx echo.Context, id plans.Id) error {
	plans, err := pc.plansUsecase.GetPlanHistory(ctx, id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, convertToPlanHistoryDto(plans))
}

func (pc *Plans) PostUpgradePlan(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "post upgrade plan")
}

func (pc *Plans) GetPlanUsage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "get plan usage")
}

func convertToCurrentPlanDto(plan *schema.UserPlans) dto.PlanDto {
	startDate := ""
	endDate := ""
	if plan.Start_Date != nil {
		startDate = plan.Start_Date.String()
	}
	if plan.End_Date != nil {
		endDate = plan.End_Date.String()
	}
	return dto.PlanDto{
		UUID:      plan.UUID.String(),
		PlanType:  string(plan.PlanType),
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func convertToPlanHistoryDto(plans []schema.UserPlans) []dto.PlanDto {
	var dto = make([]dto.PlanDto, len(plans))
	for i, v := range plans {
		startDate := ""
		endDate := ""
		if v.Start_Date != nil {
			startDate = v.Start_Date.String()
		}
		if v.End_Date != nil {
			endDate = v.End_Date.String()
		}
		dto[i].UUID = v.UUID.String()
		dto[i].StartDate = startDate
		dto[i].EndDate = endDate
		dto[i].PlanType = string(v.PlanType)
	}
	return dto
}
