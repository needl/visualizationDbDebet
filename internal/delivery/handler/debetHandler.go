package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/debet"
)

type DebetHandler struct {
	service *debet.Service
}

func NewDebetHandler(service *debet.Service) *DebetHandler {
	return &DebetHandler{service: service}
}

func (h *DebetHandler) GetAllDebet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	debets, err := h.service.GetAllDebets(ctx)
	if err != nil {
		slog.Error(err.Error(), "info", "упал на получение данных в хендлере")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondJSON(w, debets)
	slog.Info("Get all debets")
}

func (h *DebetHandler) GetByOrgName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgName := vars["orgName"]
	if orgName == "" {
		slog.Warn("OrgName is empty")
		http.Error(w, "OrgName is empty", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	deb, err := h.service.GetByOrgName(ctx, orgName)

	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if deb == nil {
		slog.Warn("Debet not found", "orgName", orgName)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	respondJSON(w, deb)
	slog.Info("Get debet by orgName", "orgName", orgName)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	buf, err := json.Marshal(data)
	if err != nil {
		slog.Error(err.Error(), "info", "Ошибка парса ответа")
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(buf); err != nil {
		slog.Error("failed to write response", "error", err)
	}
}
