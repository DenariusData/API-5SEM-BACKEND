package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type TempoHandler struct {
	useCase *usecase.TempoUseCase
}

func NewTempoHandler(uc *usecase.TempoUseCase) *TempoHandler {
	return &TempoHandler{useCase: uc}
}

func (h *TempoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch tempo"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
