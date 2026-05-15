package response

import (
	"context"
	"log/slog"
	"net/http"
	"visualizationDbDebet/internal/httpx"
)

type Handler struct {
	service service
}

type service interface {
	GetResponse(ctx context.Context) (*Response, error)
	GetResponseWithMIP(ctx context.Context) (*Response, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetResponse(w http.ResponseWriter, r *http.Request) {
	pageDTO, err := h.service.GetResponse(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, pageDTO)
	slog.Info("GetResponse", "pageDto", pageDTO)
}

func (h *Handler) GetResponseWithMIP(w http.ResponseWriter, r *http.Request) {
	pageDTO, err := h.service.GetResponseWithMIP(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, pageDTO)
	slog.Info("GetResponseWithMIP", "pageDto", pageDTO)
}
