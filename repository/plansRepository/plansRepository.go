package plansRepository

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/plans"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type plansRepository struct {
	db *gorm.DB
}

func NewPlansRepository(db *gorm.DB) *plansRepository {
	return &plansRepository{
		db: db,
	}
}
func (pr *plansRepository) GetCurrentPlan(ctx echo.Context, id plans.Id) (*schema.UserPlans, error) {
	var currentplan schema.UserPlans
	tx := pr.db.Where("user_uuid = ? AND is_active = ?", id, true).Take(&currentplan)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &currentplan, nil
}

func (pr *plansRepository) GetPlanHistory(ctx echo.Context, id plans.Id) ([]schema.UserPlans, error) {
	var planHistory []schema.UserPlans
	tx := pr.db.Where("user_uuid = ?", id).Find(&planHistory)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return planHistory, nil
}
