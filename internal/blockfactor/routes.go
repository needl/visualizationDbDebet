package blockfactor

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	blockFactorsRouter := r.PathPrefix("/block-factors").Subrouter()
	blockFactorsRouter.HandleFunc("", h.GetAll).Methods("GET")
	blockFactorsRouter.HandleFunc("/{id}", h.GetByID).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/blockFactor").Subrouter()
	legacyRouter.HandleFunc("", h.GetAll).Methods("GET")
	legacyRouter.HandleFunc("/{id}", h.GetByID).Methods("GET")
}
