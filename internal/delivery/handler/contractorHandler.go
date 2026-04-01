package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/contractor"
	"visualizationBdDebet/internal/delivery/util"
)

type ContractorHandler struct {
	service *contractor.Service
}

func NewContractorHandler(service *contractor.Service) *ContractorHandler {
	return &ContractorHandler{service: service}
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
	slog.Info("Get Contractors With BlockFactors: ", "contractors", contractors)
}
