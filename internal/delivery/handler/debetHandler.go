package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/debet"
	"visualizationBdDebet/internal/delivery/util"
)

type DebetHandler struct {
	service *debet.Service
}

func NewDebetHandler(service *debet.Service) *DebetHandler {
	return &DebetHandler{service: service}
}

func (h *DebetHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	debets, err := h.service.GetAll(ctx)
	if err != nil {
		//slog.Error(err.Error(), "info", "упал на получение данных в хендлере")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, debets)
	slog.Info("Get all debets")
}

func (h *DebetHandler) GetByOrgName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgName := vars["orgName"]
	if orgName == "" {
		//slog.Warn("OrgName is empty")
		http.Error(w, "OrgName is empty", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	view, err := h.service.GetByOrgName(ctx, orgName)

	if err != nil {
		//slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if view == nil {
		//slog.Warn("View not found", "orgName", orgName)
		http.Error(w, fmt.Sprintf("Debet with org name '%s' not found", orgName), http.StatusNotFound)
		return
	}

	util.RespondJSON(w, view)
	slog.Info("Get debet view by orgName", "orgName", orgName)
}
