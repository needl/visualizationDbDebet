package blockfactor

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	blockFactorsRouter := r.PathPrefix("/block-factors").Subrouter()
	blockFactorsRouter.HandleFunc("", h.getAll).Methods("GET")
	blockFactorsRouter.HandleFunc("/{id}", h.getByID).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/blockFactor").Subrouter()
	legacyRouter.HandleFunc("", h.getAll).Methods("GET")
	legacyRouter.HandleFunc("/{id}", h.getByID).Methods("GET")
}
