package customer

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	customerRouter := r.PathPrefix("/customer").Subrouter()
	customerRouter.HandleFunc("", h.GetCustomers).Methods("GET")
	customerRouter.HandleFunc("/summary/{orgName}", h.GetSummaryByCustomerID).Methods("GET")
	customerRouter.HandleFunc("/top-debtors/{orgName}", h.GetTopItemsByCustomerID).Methods("GET")
	customerRouter.HandleFunc("/top-debtors-overdue/{orgName}", h.GetTopItemsOverdueByCustomerID).Methods("GET")
	customerRouter.HandleFunc("/blockFactors/{orgName}", h.GetCountBlockFactorsByCustomerID).Methods("GET")
}
