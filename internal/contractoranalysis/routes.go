package contractoranalysis

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	analysisRouter := r.PathPrefix("/contractor-analytics").Subrouter()
	analysisRouter.HandleFunc("", h.getContractors).Methods("GET")
	analysisRouter.HandleFunc("/{contractorName}", h.getAnalytics).Methods("GET")
	analysisRouter.HandleFunc("/{contractorName}/object-details", h.getObjectDetails).Methods("GET")
}
