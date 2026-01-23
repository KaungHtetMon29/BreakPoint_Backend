package breakpointRepository

import (
	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/breakpoints"
	"github.com/KaungHtetMon29/BreakPoint_Backend/db/schema"
	"github.com/KaungHtetMon29/BreakPoint_Backend/internal/ai"
	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go/v3"
	"gorm.io/gorm"
)

type breakpointRepository struct {
	db     *gorm.DB
	client *openai.Client
}

func NewBreakpointRepository(db *gorm.DB, client *openai.Client) *breakpointRepository {
	return &breakpointRepository{
		db:     db,
		client: client,
	}
}

func (br *breakpointRepository) GenerateBreakPoint(ctx echo.Context) error {
	ai.Request(*br.client)
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
