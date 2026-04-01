package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DenariusData/API-5SEM-BACKEND/internal/usecase"
)

type ResponsavelHandler struct {
	useCase *usecase.ResponsavelUseCase
}

func NewResponsavelHandler(uc *usecase.ResponsavelUseCase) *ResponsavelHandler {
	return &ResponsavelHandler{useCase: uc}
}

func (h *ResponsavelHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.useCase.GetAll()
	if err != nil {
		http.Error(w, `{"error":"failed to fetch responsaveis"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
