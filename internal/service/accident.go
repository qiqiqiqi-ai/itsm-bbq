package service

import (
	"context"
	"itsm/internal/model"
	"itsm/internal/repository"
)

type AccidentService interface {
	GetAccidentList(ctx context.Context) (*model.Accident, error)
}

func NewAccidentService(
	service *Service,
	accidentRepository repository.AccidentRepository,
) AccidentService {
	return &accidentService{
		Service:            service,
		accidentRepository: accidentRepository,
	}
}

type accidentService struct {
	*Service
	accidentRepository repository.AccidentRepository
}

func (s *accidentService) GetAccidentList(ctx context.Context) (*model.Accident, error) {
	return s.accidentRepository.GetAccidentList(ctx)
}
