package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type MaterialRepository interface {
	FindAll() ([]entity.DimMaterial, error)
}

type MaterialUseCase struct {
	repo MaterialRepository
}

func NewMaterialUseCase(repo MaterialRepository) *MaterialUseCase {
	return &MaterialUseCase{repo: repo}
}

func (uc *MaterialUseCase) GetAll() ([]entity.DimMaterial, error) {
	return uc.repo.FindAll()
}
