package customer

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	customersRouter := r.PathPrefix("/customers").Subrouter()
	customersRouter.HandleFunc("", h.GetCustomers).Methods("GET")
	customersRouter.HandleFunc("/summary/{orgName}", h.GetSummaryByCustomerID).Methods("GET")
	customersRouter.HandleFunc("/top-debtors/{orgName}", h.GetTopItemsByCustomerID).Methods("GET")
	customersRouter.HandleFunc("/top-debtors-overdue/{orgName}", h.GetTopItemsOverdueByCustomerID).Methods("GET")
	customersRouter.HandleFunc("/block-factors/{orgName}", h.GetCountBlockFactorsByCustomerID).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/customer").Subrouter()
	legacyRouter.HandleFunc("", h.GetCustomers).Methods("GET")
	legacyRouter.HandleFunc("/summary/{orgName}", h.GetSummaryByCustomerID).Methods("GET")
	legacyRouter.HandleFunc("/top-debtors/{orgName}", h.GetTopItemsByCustomerID).Methods("GET")
	legacyRouter.HandleFunc("/top-debtors-overdue/{orgName}", h.GetTopItemsOverdueByCustomerID).Methods("GET")
	legacyRouter.HandleFunc("/blockFactors/{orgName}", h.GetCountBlockFactorsByCustomerID).Methods("GET")
}
