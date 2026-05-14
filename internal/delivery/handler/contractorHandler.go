package handler

import (
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/contractor"
	"visualizationBdDebet/internal/delivery/util"

	"github.com/gorilla/mux"
)

type ContractorHandler struct {
	service *contractor.Service
}

func NewHandlerContractor(service *contractor.Service) *ContractorHandler {
	return &ContractorHandler{service: service}
}

func (h *ContractorHandler) GetContractorsWithCurrDebet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithCurrDeb(ctx)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get contractors with current debet", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorsWithOverdueDebet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithOverdueDeb(ctx)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get contractors with overdue debet", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorsWithBlockFactors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["orgName"]
	columnName := r.URL.Query().Get("columnName")

	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithBlockFactors(ctx, sourceOrgName, columnName)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get contractors with block factors", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorsWithDebt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["orgName"]
	counterpartyName := r.URL.Query().Get("counterpartyName")

	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithDebt(ctx, sourceOrgName, counterpartyName)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get contractors for debt", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorsWithOverdue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["orgName"]
	counterpartyName := r.URL.Query().Get("counterpartyName")

	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithOverdue(ctx, sourceOrgName, counterpartyName)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get contractors for overdue", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorForTable(w http.ResponseWriter, r *http.Request) {
	counterpartyName := r.URL.Query().Get("counterpartyName")

	ctx := r.Context()
	contractors, err := h.service.GetContractorForTable(ctx, counterpartyName)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get contractors for table", "contractors", contractors)
}
