package contractor

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	contractorsRouter := r.PathPrefix("/contractors").Subrouter()
	contractorsRouter.HandleFunc("/table", h.GetContractorForTable).Methods("GET")
	contractorsRouter.HandleFunc("/debts/current", h.GetContractorsWithCurrDebet).Methods("GET")
	contractorsRouter.HandleFunc("/debts/overdue", h.GetContractorsWithOverdueDebet).Methods("GET")
	contractorsRouter.HandleFunc("/{orgName}/debts", h.GetContractorsWithDebt).Methods("GET")
	contractorsRouter.HandleFunc("/{orgName}/overdue-debts", h.GetContractorsWithOverdue).Methods("GET")
	contractorsRouter.HandleFunc("/{orgName}", h.GetContractorsWithBlockFactors).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/contractor").Subrouter()
	legacyRouter.HandleFunc("/table", h.GetContractorForTable).Methods("GET")
	legacyRouter.HandleFunc("/debet/curr", h.GetContractorsWithCurrDebet).Methods("GET")
	legacyRouter.HandleFunc("/debet/overdue", h.GetContractorsWithOverdueDebet).Methods("GET")
	legacyRouter.HandleFunc("/{orgName}/debt", h.GetContractorsWithDebt).Methods("GET")
	legacyRouter.HandleFunc("/{orgName}/overdue", h.GetContractorsWithOverdue).Methods("GET")
	legacyRouter.HandleFunc("/{orgName}", h.GetContractorsWithBlockFactors).Methods("GET")
}
