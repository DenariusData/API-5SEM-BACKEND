package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type TarefaHandler struct {
	useCase *usecase.TarefaUseCase
}

func NewTarefaHandler(uc *usecase.TarefaUseCase) *TarefaHandler {
	return &TarefaHandler{useCase: uc}
}

func (h *TarefaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch tarefas"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
