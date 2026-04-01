package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type FatoComprasHandler struct {
	useCase *usecase.FatoComprasUseCase
}

func NewFatoComprasHandler(uc *usecase.FatoComprasUseCase) *FatoComprasHandler {
	return &FatoComprasHandler{useCase: uc}
}

func (h *FatoComprasHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch fato compras"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
