package blockfactor

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
	factors, err := h.service.GetAllView(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, factors)
	slog.Info("Get all factors")
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	view, err := h.service.GetViewById(r.Context(), id)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, view)
	slog.Info("Get blockfactor by id", "id", id)
}
