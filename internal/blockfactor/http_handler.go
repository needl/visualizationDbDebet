package blockfactor

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
	getAllView(ctx context.Context) ([]View, error)
	getViewByID(ctx context.Context, id string) (*View, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	factors, err := h.service.getAllView(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, factors)
	slog.Info("Get all factors")
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	view, err := h.service.getViewByID(r.Context(), id)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, view)
	slog.Info("Get blockfactor by id", "id", id)
}
