package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

const (
	ContentTypeHeader = "Content-Type"
	ContentTypeJSON   = "application/json"
)

type ProjetoHandler struct {
	useCase *usecase.ProjetoUseCase
}

func NewProjetoHandler(uc *usecase.ProjetoUseCase) *ProjetoHandler {
	return &ProjetoHandler{useCase: uc}
}

func (h *ProjetoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	projects, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch projects"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjetoHandler) GetInvestimentoByPrograma(w http.ResponseWriter, r *http.Request) {
	investimentos, err := h.useCase.GetInvestimentoByPrograma()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch investment by program"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	json.NewEncoder(w).Encode(investimentos)
}

func (h *ProjetoHandler) GetMateriaisByProjeto(w http.ResponseWriter, r *http.Request) {
	materiais, err := h.useCase.GetMateriaisByProjeto()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch materials by project"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeader, ContentTypeJSON)
	json.NewEncoder(w).Encode(materiais)
}
