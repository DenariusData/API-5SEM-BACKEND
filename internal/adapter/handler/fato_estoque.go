package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type FatoEstoqueHandler struct {
	useCase *usecase.FatoEstoqueUseCase
}

func NewFatoEstoqueHandler(uc *usecase.FatoEstoqueUseCase) *FatoEstoqueHandler {
	return &FatoEstoqueHandler{useCase: uc}
}

func (h *FatoEstoqueHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch fato estoque"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
