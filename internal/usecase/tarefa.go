package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type TarefaRepository interface {
	FindAll() ([]entity.DimTarefa, error)
}

type TarefaUseCase struct {
	repo TarefaRepository
}

func NewTarefaUseCase(repo TarefaRepository) *TarefaUseCase {
	return &TarefaUseCase{repo: repo}
}

func (uc *TarefaUseCase) GetAll() ([]entity.DimTarefa, error) {
	return uc.repo.FindAll()
}
