package contractor

import (
	"log/slog"
	"net/http"
	"visualizationDbDebet/internal/httpx"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getContractorsWithCurrDebet(w http.ResponseWriter, r *http.Request) {
	contractors, err := h.service.getContractorsWithCurrDeb(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, contractors)
	slog.Info("Get contractors with current debet", "contractors", contractors)
}

func (h *Handler) getContractorsWithOverdueDebet(w http.ResponseWriter, r *http.Request) {
	contractors, err := h.service.getContractorsWithOverdueDeb(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, contractors)
	slog.Info("Get contractors with overdue debet", "contractors", contractors)
}

func (h *Handler) getContractorsWithBlockFactors(w http.ResponseWriter, r *http.Request) {
	sourceOrgName := mux.Vars(r)["orgName"]
	columnName := r.URL.Query().Get("columnName")

	contractors, err := h.service.getContractorsWithBlockFactors(r.Context(), sourceOrgName, columnName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, contractors)
	slog.Info("Get contractors with block factors", "contractors", contractors)
}

func (h *Handler) getContractorsWithDebt(w http.ResponseWriter, r *http.Request) {
	sourceOrgName := mux.Vars(r)["orgName"]
	counterpartyName := r.URL.Query().Get("counterpartyName")

	contractors, err := h.service.getContractorsWithDebt(r.Context(), sourceOrgName, counterpartyName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, contractors)
	slog.Info("Get contractors for debt", "contractors", contractors)
}

func (h *Handler) getContractorsWithOverdue(w http.ResponseWriter, r *http.Request) {
	sourceOrgName := mux.Vars(r)["orgName"]
	counterpartyName := r.URL.Query().Get("counterpartyName")

	contractors, err := h.service.getContractorsWithOverdue(r.Context(), sourceOrgName, counterpartyName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, contractors)
	slog.Info("Get contractors for overdue", "contractors", contractors)
}

func (h *Handler) getContractorForTable(w http.ResponseWriter, r *http.Request) {
	counterpartyName := r.URL.Query().Get("counterpartyName")

	contractors, err := h.service.getContractorForTable(r.Context(), counterpartyName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, contractors)
	slog.Info("Get contractors for table", "contractors", contractors)
}
