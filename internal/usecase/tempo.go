package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type TempoRepository interface {
	FindAll() ([]entity.DimTempo, error)
}

type TempoUseCase struct {
	repo TempoRepository
}

func NewTempoUseCase(repo TempoRepository) *TempoUseCase {
	return &TempoUseCase{repo: repo}
}

func (uc *TempoUseCase) GetAll() ([]entity.DimTempo, error) {
	return uc.repo.FindAll()
}
