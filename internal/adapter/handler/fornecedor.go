package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type FornecedorHandler struct {
	useCase *usecase.FornecedorUseCase
}

func NewFornecedorHandler(uc *usecase.FornecedorUseCase) *FornecedorHandler {
	return &FornecedorHandler{useCase: uc}
}

func (h *FornecedorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch fornecedores"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
