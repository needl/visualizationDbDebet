package debet

import (
	"context"
	"log/slog"
	"net/http"
	"visualizationDbDebet/internal/httpx"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service
}

type service interface {
	GetAll(ctx context.Context) ([]View, error)
	GetAllWithMIP(ctx context.Context) ([]View, error)
	GetByOrgName(ctx context.Context, orgName string) (*View, error)
}

func NewHandler(service service) *Handler {
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
