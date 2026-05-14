package handler

import (
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/blockfactor"
	"visualizationBdDebet/internal/delivery/util"

	"github.com/gorilla/mux"
)

type BlockFactorHandler struct {
	service *blockfactor.Service
}

func NewHandlerBlockFactor(service *blockfactor.Service) *BlockFactorHandler {
	return &BlockFactorHandler{service: service}
}

func (h *BlockFactorHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	factors, err := h.service.GetAllView(ctx)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, factors)
	slog.Info("Get all factors")
}

func (h *BlockFactorHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	view, err := h.service.GetViewById(ctx, id)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, view)
	slog.Info("Get blockfactor by id", "id", id)
}
