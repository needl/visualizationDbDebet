package response

import (
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/delivery/util"
)

type Handler struct {
	service *Service
}

func NewResponseHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetResponse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageDto, err := h.service.GetResponse(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, pageDto)
	slog.Info("GetResponse: ", "pageDto", pageDto)
}

func (h *Handler) GetResponseWithMIP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageDto, err := h.service.GetResponseWithMIP(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, pageDto)
	slog.Info("GetResponse: ", "pageDto", pageDto)
}
