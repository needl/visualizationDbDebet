package customer

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	customersRouter := r.PathPrefix("/customers").Subrouter()
	customersRouter.HandleFunc("", h.getCustomers).Methods("GET")
	customersRouter.HandleFunc("/summary/{orgName}", h.getSummaryByCustomerID).Methods("GET")
	customersRouter.HandleFunc("/top-debtors/{orgName}", h.getTopItemsByCustomerID).Methods("GET")
	customersRouter.HandleFunc("/top-debtors-overdue/{orgName}", h.getTopItemsOverdueByCustomerID).Methods("GET")
	customersRouter.HandleFunc("/block-factors/{orgName}", h.getCountBlockFactorsByCustomerID).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/customer").Subrouter()
	legacyRouter.HandleFunc("", h.getCustomers).Methods("GET")
	legacyRouter.HandleFunc("/summary/{orgName}", h.getSummaryByCustomerID).Methods("GET")
	legacyRouter.HandleFunc("/top-debtors/{orgName}", h.getTopItemsByCustomerID).Methods("GET")
	legacyRouter.HandleFunc("/top-debtors-overdue/{orgName}", h.getTopItemsOverdueByCustomerID).Methods("GET")
	legacyRouter.HandleFunc("/blockFactors/{orgName}", h.getCountBlockFactorsByCustomerID).Methods("GET")
}
