package usecase

import (
	"errors"
	"testing"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/domain/entity"
)

// Mock
type mockProjetoRepo struct {	
	projetos []entity.DimProjeto
	err      error
}

func (m *mockProjetoRepo) FindAll() ([]entity.DimProjeto, error) {
	return m.projetos, m.err
}

// Teste 1: sucesso
func TestProjetoUseCase_GetAll_Success(t *testing.T) {
	fake := []entity.DimProjeto{
		{IdProjeto: "1", NomeProjeto: "Projeto Alpha", Status: "Ativo"},
		{IdProjeto: "2", NomeProjeto: "Projeto Beta", Status: "Concluído"},
	}
	mock := &mockProjetoRepo{projetos: fake, err: nil}
	uc := NewProjetoUseCase(mock)

	result, err := uc.GetAll()

	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("esperava 2 projetos, recebeu: %d", len(result))
	}
}

// Teste 2: erro do banco
func TestProjetoUseCase_GetAll_Error(t *testing.T) {
	mock := &mockProjetoRepo{
		projetos: nil,
		err:      errors.New("connection refused"),
	}
	uc := NewProjetoUseCase(mock)

	result, err := uc.GetAll()

	if err == nil {
		t.Fatal("esperava erro, recebeu nil")
	}
	if result != nil {
		t.Errorf("esperava nil, recebeu: %v", result)
	}
}

// Teste 3: lista vazia
func TestProjetoUseCase_GetAll_Empty(t *testing.T) {
	mock := &mockProjetoRepo{projetos: []entity.DimProjeto{}, err: nil}
	uc := NewProjetoUseCase(mock)

	result, err := uc.GetAll()

	if err != nil {
		t.Fatalf("esperava sem erro, recebeu: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("esperava 0 projetos, recebeu: %d", len(result))
	}
}