package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/contract"
	"visualizationBdDebet/internal/delivery/util"
)

type ContractHandler struct {
	service *contract.Service
}

func NewContractHandler(service *contract.Service) *ContractHandler {
	return &ContractHandler{service: service}
}

// GetAll возвращает все записи View contracts
func (h *ContractHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contracts, err := h.service.GetAll(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	util.RespondJSON(w, contracts)
	slog.Info("Get all contracts")
}

// GetById возвращается View из записи таблицы contracts по id
func (h *ContractHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "Invalid value", http.StatusBadRequest)
	}

	ctx := r.Context()
	view, err := h.service.GetById(ctx, id)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if view == nil {
		http.Error(w, fmt.Sprintf("Contract with id '%s' not found", id), http.StatusNotFound)
	}

	util.RespondJSON(w, view)
	slog.Info("Get contract view by id", "id", id)
}
