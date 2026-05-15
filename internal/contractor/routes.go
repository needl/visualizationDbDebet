package contractor

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	contractorRouter := r.PathPrefix("/contractor").Subrouter()
	contractorRouter.HandleFunc("/table", h.GetContractorForTable).Methods("GET")
	contractorRouter.HandleFunc("/debet/curr", h.GetContractorsWithCurrDebet).Methods("GET")
	contractorRouter.HandleFunc("/debet/overdue", h.GetContractorsWithOverdueDebet).Methods("GET")
	contractorRouter.HandleFunc("/{orgName}/debt", h.GetContractorsWithDebt).Methods("GET")
	contractorRouter.HandleFunc("/{orgName}/overdue", h.GetContractorsWithOverdue).Methods("GET")
	contractorRouter.HandleFunc("/{orgName}", h.GetContractorsWithBlockFactors).Methods("GET")
}
