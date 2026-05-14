package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/debet"
	"visualizationBdDebet/internal/delivery/util"

	"github.com/gorilla/mux"
)

type DebetHandler struct {
	service *debet.Service
}

func NewHandlerDebet(service *debet.Service) *DebetHandler {
	return &DebetHandler{service: service}
}

func (h *DebetHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	debets, err := h.service.GetAll(ctx)
	if err != nil {
		util.RespondError(w, err, "Internal server error")
		return
	}

	util.RespondJSON(w, debets)
	slog.Info("Get all debets")
}

func (h *DebetHandler) GetAllWithMIP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	debets, err := h.service.GetAllWithMIP(ctx)
	if err != nil {
		util.RespondError(w, err, "Internal server error")
		return
	}

	util.RespondJSON(w, debets)
	slog.Info("Get all debets with MIP")
}

func (h *DebetHandler) GetByOrgName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgName := vars["orgName"]
	if orgName == "" {
		http.Error(w, "OrgName is empty", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	view, err := h.service.GetByOrgName(ctx, orgName)
	if err != nil {
		util.RespondError(w, err, "Internal server error")
		return
	}

	if view == nil {
		http.Error(w, fmt.Sprintf("Debet with org name '%s' not found", orgName), http.StatusNotFound)
		return
	}

	util.RespondJSON(w, view)
	slog.Info("Get debet view by orgName", "orgName", orgName)
}
