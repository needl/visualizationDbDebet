package contract

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	contractsRouter := r.PathPrefix("/contracts").Subrouter()
	contractsRouter.HandleFunc("", h.GetAll).Methods("GET")
	contractsRouter.HandleFunc("/{id}", h.GetByID).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/contract").Subrouter()
	legacyRouter.HandleFunc("", h.GetAll).Methods("GET")
	legacyRouter.HandleFunc("/{id}", h.GetByID).Methods("GET")
}
