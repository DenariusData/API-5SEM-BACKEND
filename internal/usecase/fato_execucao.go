package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type FatoExecucaoRepository interface {
	FindAll() ([]entity.FatoExecucaoTarefas, error)
}

type FatoExecucaoUseCase struct {
	repo FatoExecucaoRepository
}

func NewFatoExecucaoUseCase(repo FatoExecucaoRepository) *FatoExecucaoUseCase {
	return &FatoExecucaoUseCase{repo: repo}
}

func (uc *FatoExecucaoUseCase) GetAll() ([]entity.FatoExecucaoTarefas, error) {
	return uc.repo.FindAll()
}
