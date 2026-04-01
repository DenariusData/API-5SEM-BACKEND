package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type SolicitacaoHandler struct {
	useCase *usecase.SolicitacaoUseCase
}

func NewSolicitacaoHandler(uc *usecase.SolicitacaoUseCase) *SolicitacaoHandler {
	return &SolicitacaoHandler{useCase: uc}
}

func (h *SolicitacaoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch solicitacoes"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
