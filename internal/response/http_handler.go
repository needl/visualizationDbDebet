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
	getResponse(ctx context.Context) (*Response, error)
	getResponseWithMIP(ctx context.Context) (*Response, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getResponse(w http.ResponseWriter, r *http.Request) {
	pageDTO, err := h.service.getResponse(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, pageDTO)
	slog.Info("GetResponse", "pageDto", pageDTO)
}

func (h *Handler) getResponseWithMIP(w http.ResponseWriter, r *http.Request) {
	pageDTO, err := h.service.getResponseWithMIP(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, pageDTO)
	slog.Info("GetResponseWithMIP", "pageDto", pageDTO)
}
