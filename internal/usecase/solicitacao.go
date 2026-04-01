package usecase

import "github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"

type SolicitacaoRepository interface {
	FindAll() ([]entity.DimSolicitacao, error)
}

type SolicitacaoUseCase struct {
	repo SolicitacaoRepository
}

func NewSolicitacaoUseCase(repo SolicitacaoRepository) *SolicitacaoUseCase {
	return &SolicitacaoUseCase{repo: repo}
}

func (uc *SolicitacaoUseCase) GetAll() ([]entity.DimSolicitacao, error) {
	return uc.repo.FindAll()
}
