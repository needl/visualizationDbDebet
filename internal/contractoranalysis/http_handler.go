package contractoranalysis

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
	getContractors(ctx context.Context) ([]string, error)
	getAnalytics(ctx context.Context, contractorName string) (*Analytics, error)
	getObjectDetails(ctx context.Context, contractorName string, customerName string, objectName string) (*ObjectDetails, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getContractors(w http.ResponseWriter, r *http.Request) {
	contractors, err := h.service.getContractors(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, contractors)
	slog.Info("GetContractors from contractoranalyses")
}

func (h *Handler) getAnalytics(w http.ResponseWriter, r *http.Request) {
	contractorName := mux.Vars(r)["contractorName"]

	analytics, err := h.service.getAnalytics(r.Context(), contractorName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, analytics)
	slog.Info("GetAnalytics from contractoranalyses")
}

func (h *Handler) getObjectDetails(w http.ResponseWriter, r *http.Request) {
	contractorName := mux.Vars(r)["contractorName"]
	customerName := r.URL.Query().Get("customerName")
	objectName := r.URL.Query().Get("objectName")

	details, err := h.service.getObjectDetails(r.Context(), contractorName, customerName, objectName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, details)
	slog.Info("GetObjectDetails from contractoranalyses")
}
