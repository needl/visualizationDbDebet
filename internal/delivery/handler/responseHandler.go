package handler

import (
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/delivery/util"
	"visualizationBdDebet/internal/response"
)

type ResponseHandler struct {
	service *response.Service
}

func NewResponseHandler(service *response.Service) *ResponseHandler {
	return &ResponseHandler{service: service}
}

func (h *ResponseHandler) GetResponse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageDto, err := h.service.GetResponse(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, pageDto)
	slog.Info("GetResponse: ", "pageDto", pageDto)
}

func (h *ResponseHandler) GetResponseWithMIP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageDto, err := h.service.GetResponseWithMIP(ctx)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, pageDto)
	slog.Info("GetResponse: ", "pageDto", pageDto)
}
