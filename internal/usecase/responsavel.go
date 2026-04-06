package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type ResponsavelRepository interface {
	FindAll() ([]entity.DimResponsavel, error)
}

type ResponsavelUseCase struct {
	repo ResponsavelRepository
}

func NewResponsavelUseCase(repo ResponsavelRepository) *ResponsavelUseCase {
	return &ResponsavelUseCase{repo: repo}
}

func (uc *ResponsavelUseCase) GetAll() ([]entity.DimResponsavel, error) {
	return uc.repo.FindAll()
}
