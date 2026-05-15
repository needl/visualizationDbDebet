package contract

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	contractRouter := r.PathPrefix("/contract").Subrouter()
	contractRouter.HandleFunc("", h.GetAll).Methods("GET")
	contractRouter.HandleFunc("/{id}", h.GetByID).Methods("GET")
}
