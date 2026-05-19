package contract

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
	getByID(ctx context.Context, id string) (*View, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	contracts, err := h.service.getAll(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, contracts)
	slog.Info("Get all contracts")
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	view, err := h.service.getByID(r.Context(), id)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, view)
	slog.Info("Get contract view by id", "id", id)
}
