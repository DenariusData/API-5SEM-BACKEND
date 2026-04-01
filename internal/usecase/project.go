package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type ProjetoRepository interface {
	FindAll() ([]entity.DimProjeto, error)
}

type ProjetoUseCase struct {
	repo ProjetoRepository
}

func NewProjetoUseCase(repo ProjetoRepository) *ProjetoUseCase {
	return &ProjetoUseCase{repo: repo}
}

func (uc *ProjetoUseCase) GetAll() ([]entity.DimProjeto, error) {
	return uc.repo.FindAll()
}
