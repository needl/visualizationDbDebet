package contractor

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	contractorsRouter := r.PathPrefix("/contractors").Subrouter()
	contractorsRouter.HandleFunc("/table", h.getContractorForTable).Methods("GET")
	contractorsRouter.HandleFunc("/debts/current", h.getContractorsWithCurrDebet).Methods("GET")
	contractorsRouter.HandleFunc("/debts/overdue", h.getContractorsWithOverdueDebet).Methods("GET")
	contractorsRouter.HandleFunc("/{orgName}/debts", h.getContractorsWithDebt).Methods("GET")
	contractorsRouter.HandleFunc("/{orgName}/overdue-debts", h.getContractorsWithOverdue).Methods("GET")
	contractorsRouter.HandleFunc("/{orgName}", h.getContractorsWithBlockFactors).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/contractor").Subrouter()
	legacyRouter.HandleFunc("/table", h.getContractorForTable).Methods("GET")
	legacyRouter.HandleFunc("/debet/curr", h.getContractorsWithCurrDebet).Methods("GET")
	legacyRouter.HandleFunc("/debet/overdue", h.getContractorsWithOverdueDebet).Methods("GET")
	legacyRouter.HandleFunc("/{orgName}/debt", h.getContractorsWithDebt).Methods("GET")
	legacyRouter.HandleFunc("/{orgName}/overdue", h.getContractorsWithOverdue).Methods("GET")
	legacyRouter.HandleFunc("/{orgName}", h.getContractorsWithBlockFactors).Methods("GET")
}
