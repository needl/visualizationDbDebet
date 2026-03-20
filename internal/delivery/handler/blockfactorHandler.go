package handler

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/blockfactor"
	"visualizationBdDebet/internal/delivery/util"
)

type BlockFactorHandler struct {
	service *blockfactor.Service
}

func NewBlockFactorHandler(service *blockfactor.Service) *BlockFactorHandler {
	return &BlockFactorHandler{service: service}
}

func (h *BlockFactorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	factors, err := h.service.GetAllView(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	util.RespondJSON(w, factors)
	slog.Info("Get all factors")
}

func (h *BlockFactorHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "Invalid value", http.StatusBadRequest)
	}

	ctx := r.Context()
	view, err := h.service.GetViewById(ctx, id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if view == nil {
		http.Error(w, "View not found", http.StatusNotFound)
	}

	util.RespondJSON(w, view)
	slog.Info("Get blockfactor by id", "id", id)
}
