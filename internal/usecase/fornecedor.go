package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type FornecedorRepository interface {
	FindAll() ([]entity.DimFornecedor, error)
}

type FornecedorUseCase struct {
	repo FornecedorRepository
}

func NewFornecedorUseCase(repo FornecedorRepository) *FornecedorUseCase {
	return &FornecedorUseCase{repo: repo}
}

func (uc *FornecedorUseCase) GetAll() ([]entity.DimFornecedor, error) {
	return uc.repo.FindAll()
}
