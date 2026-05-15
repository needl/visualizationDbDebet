package debet

import (
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/httpx"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	debets, err := h.service.GetAll(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, debets)
	slog.Info("Get all debets")
}

func (h *Handler) GetAllWithMIP(w http.ResponseWriter, r *http.Request) {
	debets, err := h.service.GetAllWithMIP(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, debets)
	slog.Info("Get all debets with MIP")
}

func (h *Handler) GetByOrgName(w http.ResponseWriter, r *http.Request) {
	orgName := mux.Vars(r)["orgName"]

	view, err := h.service.GetByOrgName(r.Context(), orgName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, view)
	slog.Info("Get debet view by orgName", "orgName", orgName)
}
