package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type ProjetoRepository interface {
	FindAll() ([]entity.DimProjeto, error)
	FindInvestimentoByPrograma() ([]entity.ProgramaInvestimento, error)
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

func (uc *ProjetoUseCase) GetInvestimentoByPrograma() ([]entity.ProgramaInvestimento, error) {
	return uc.repo.FindInvestimentoByPrograma()
}
