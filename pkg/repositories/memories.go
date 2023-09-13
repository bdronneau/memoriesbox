package repositories

import (
	"context"
	"time"

	dbModels "github.com/bdronneau/memoriesbox/pkg/db/models"
	"github.com/bdronneau/memoriesbox/pkg/repositories/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *app) CountMemories() (int64, error) {
	a.logger.Debug("Count memories")

	count, err := dbModels.Memories().Count(context.Background(), a.DB)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (a *app) GetRandomMemories() (models.Memory, error) {
	a.logger.Debug("Retrieve random memory")

	memories, err := dbModels.Memories(qm.OrderBy("RANDOM()")).One(context.Background(), a.DB)
	if err != nil {
		return models.Memory{}, err
	}

	return models.Memory{
		ID:      memories.ID,
		Author:  memories.Author,
		Content: memories.Content,
		Append:  memories.Append.Format(time.DateOnly),
	}, nil
}

func (a *app) AddMemory(quote string, author string, date time.Time) error {
	a.logger.Debug("Add memory")

	memory := dbModels.Memory{
		Author:  author,
		Content: quote,
		Append:  date,
	}

	err := memory.Insert(context.Background(), a.DB, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
