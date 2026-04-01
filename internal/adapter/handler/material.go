package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type MaterialHandler struct {
	useCase *usecase.MaterialUseCase
}

func NewMaterialHandler(uc *usecase.MaterialUseCase) *MaterialHandler {
	return &MaterialHandler{useCase: uc}
}

func (h *MaterialHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch materiais"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
