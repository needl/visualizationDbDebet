package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/contractor"
	"visualizationBdDebet/internal/delivery/util"

	"github.com/gorilla/mux"
)

type ContractorHandler struct {
	service *contractor.Service
}

func NewContractorHandler(service *contractor.Service) *ContractorHandler {
	return &ContractorHandler{service: service}
}

/*func (h *ContractorHandler) GetContractorsWithBlockFactors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["orgName"]
	columnName := r.URL.Query().Get("columnName")

	if sourceOrgName == "" {
		http.Error(w, "Source organization name is null", http.StatusBadRequest)
		return
	}

	if columnName == "" {
		http.Error(w, "Column name is null", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithBlockFactors(ctx, sourceOrgName, columnName)
	if err != nil {
		http.Error(w,
			fmt.Sprintf(
				"Contractors with orgName '%s' or columnName '%s' not found",
				sourceOrgName,
				columnName,
			),
			http.StatusNotFound,
		)
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get Contractors With Curr debet: ", "contractors", contractors)
}*/

func (h *ContractorHandler) GetContractorsWithCurrDebet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithCurrDeb(ctx)
	if err != nil {
		http.Error(w, "Contractors with Curr debet not found", http.StatusNotFound)
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get Contractors With Overdue debet: ", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorsWithOverdueDebet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithOverdueDeb(ctx)
	if err != nil {
		http.Error(w, "Contractors with Overdue debet not found", http.StatusNotFound)
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get Contractors With BlockFactors: ", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorsWithBlockFactors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["orgName"]
	columnName := r.URL.Query().Get("columnName")

	if sourceOrgName == "" {
		http.Error(w, "Source organization name is null", http.StatusBadRequest)
		return
	}

	if columnName == "" {
		http.Error(w, "Column name is null", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithBlockFactors(ctx, sourceOrgName, columnName)
	if err != nil {
		http.Error(w,
			fmt.Sprintf(
				"Contractors with orgName '%s' or columnName '%s' not found",
				sourceOrgName,
				columnName,
			),
			http.StatusNotFound,
		)
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get Contractors With Curr debet: ", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorsWithDebt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["orgName"]
	counterpartyName := r.URL.Query().Get("counterpartyName")

	if sourceOrgName == "" {
		http.Error(w, "Source organization name is null", http.StatusBadRequest)
		return
	}

	if counterpartyName == "" {
		http.Error(w, "counterparty name is null", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithDebt(ctx, sourceOrgName, counterpartyName)
	if err != nil {
		http.Error(w,
			fmt.Sprintf(
				"Contractors with orgName '%s' or counterparty name '%s' not found",
				sourceOrgName,
				counterpartyName,
			),
			http.StatusNotFound,
		)
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get Contractors for debt:", "contractors", contractors)
}

func (h *ContractorHandler) GetContractorsWithOverdue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["orgName"]
	counterpartyName := r.URL.Query().Get("counterpartyName")

	if sourceOrgName == "" {
		http.Error(w, "Source organization name is null", http.StatusBadRequest)
		return
	}

	if counterpartyName == "" {
		http.Error(w, "counterparty name is null", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	contractors, err := h.service.GetContractorsWithOverdue(ctx, sourceOrgName, counterpartyName)
	if err != nil {
		http.Error(w,
			fmt.Sprintf(
				"Contractors with orgName '%s' or counterparty name '%s' not found",
				sourceOrgName,
				counterpartyName,
			),
			http.StatusNotFound,
		)
		return
	}

	util.RespondJSON(w, contractors)
	slog.Info("Get Contractors for overdue:", "contractors", contractors)
}
