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
	getAll(ctx context.Context) ([]View, error)
	getAllWithMIP(ctx context.Context) ([]View, error)
	getByOrgName(ctx context.Context, orgName string) (*View, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	debets, err := h.service.getAll(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, debets)
	slog.Info("Get all debets")
}

func (h *Handler) getAllWithMIP(w http.ResponseWriter, r *http.Request) {
	debets, err := h.service.getAllWithMIP(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, debets)
	slog.Info("Get all debets with MIP")
}

func (h *Handler) getByOrgName(w http.ResponseWriter, r *http.Request) {
	orgName := mux.Vars(r)["orgName"]

	view, err := h.service.getByOrgName(r.Context(), orgName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, view)
	slog.Info("Get debet view by orgName", "orgName", orgName)
}
