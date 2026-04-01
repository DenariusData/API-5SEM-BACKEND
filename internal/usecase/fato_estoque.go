package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type FatoEstoqueRepository interface {
	FindAll() ([]entity.FatoEstoqueMateriais, error)
}

type FatoEstoqueUseCase struct {
	repo FatoEstoqueRepository
}

func NewFatoEstoqueUseCase(repo FatoEstoqueRepository) *FatoEstoqueUseCase {
	return &FatoEstoqueUseCase{repo: repo}
}

func (uc *FatoEstoqueUseCase) GetAll() ([]entity.FatoEstoqueMateriais, error) {
	return uc.repo.FindAll()
}
