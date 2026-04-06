package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type FatoExecucaoHandler struct {
	useCase *usecase.FatoExecucaoUseCase
}

func NewFatoExecucaoHandler(uc *usecase.FatoExecucaoUseCase) *FatoExecucaoHandler {
	return &FatoExecucaoHandler{useCase: uc}
}

func (h *FatoExecucaoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch fato execucao"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
