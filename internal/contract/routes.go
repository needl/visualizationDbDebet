package contract

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	contractsRouter := r.PathPrefix("/contracts").Subrouter()
	contractsRouter.HandleFunc("", h.getAll).Methods("GET")
	contractsRouter.HandleFunc("/{id}", h.getByID).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/contract").Subrouter()
	legacyRouter.HandleFunc("", h.getAll).Methods("GET")
	legacyRouter.HandleFunc("/{id}", h.getByID).Methods("GET")
}
