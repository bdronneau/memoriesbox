package repositories

import (
	"context"

	dbModels "memoriesbox/pkg/db/models"
	"memoriesbox/pkg/repositories/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *app) CountMemories() int64 {
	a.logger.Debug("Count memories")

	count, err := dbModels.Memories().Count(context.Background(), a.DB)
	if err != nil {
		a.logger.Fatal(err)
	}

	return count
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
		Append:  memories.Append.Format("2006-01-02"),
	}, nil
}
