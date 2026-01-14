package breakpointRepository

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/breakpoints"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type breakpointRepository struct {
	db *gorm.DB
}

func NewBreakpointRepository(db *gorm.DB) *breakpointRepository {
	return &breakpointRepository{
		db: db,
	}
}
func (br *breakpointRepository) GenerateBreakPoint(ctx echo.Context) error {
	return nil
}

func (br *breakpointRepository) GetBreakPointTechniques(ctx echo.Context, id breakpoints.Id) ([]schema.BreakPointTechniques, error) {
	var bptechniques []schema.BreakPointTechniques
	tx := br.db.Where("user_uuid = ?", id).Find(&bptechniques)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bptechniques, nil
}

func (br *breakpointRepository) GetBreakPointHistory(ctx echo.Context, id breakpoints.Id) ([]schema.BreakPointGenerateHistory, error) {
	var bpgenhistory []schema.BreakPointGenerateHistory
	tx := br.db.Where("user_uuid = ?", id).Find(&bpgenhistory)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return bpgenhistory, nil
}
