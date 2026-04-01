package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type FatoComprasRepository interface {
	FindAll() ([]entity.FatoCompras, error)
}

type FatoComprasUseCase struct {
	repo FatoComprasRepository
}

func NewFatoComprasUseCase(repo FatoComprasRepository) *FatoComprasUseCase {
	return &FatoComprasUseCase{repo: repo}
}

func (uc *FatoComprasUseCase) GetAll() ([]entity.FatoCompras, error) {
	return uc.repo.FindAll()
}
