package handler

import (
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/contract"
	"visualizationBdDebet/internal/delivery/util"

	"github.com/gorilla/mux"
)

type ContractHandler struct {
	service *contract.Service
}

func NewHandlerContract(service *contract.Service) *ContractHandler {
	return &ContractHandler{service: service}
}

func (h *ContractHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	contracts, err := h.service.GetAll(ctx)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, contracts)
	slog.Info("Get all contracts")
}

func (h *ContractHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	view, err := h.service.GetById(ctx, id)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, view)
	slog.Info("Get contract view by id", "id", id)
}
