package repository

import (
	"context"
	"itsm/internal/model"
)

type AccidentRepository interface {
	GetAccidentList(ctx context.Context) (*model.Accident, error)
}

func NewAccidentRepository(
	repository *Repository,
) AccidentRepository {
	return &accidentRepository{
		Repository: repository,
	}
}

type accidentRepository struct {
	*Repository
}

func (r *accidentRepository) GetAccidentList(ctx context.Context) (*model.Accident, error) {
	var accident model.Accident

	return &accident, nil
}
