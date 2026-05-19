package debet

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	debetsRouter := r.PathPrefix("/debets").Subrouter()
	debetsRouter.HandleFunc("", h.getAll).Methods("GET")
	debetsRouter.HandleFunc("/with-mip", h.getAllWithMIP).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/debet").Subrouter()
	legacyRouter.HandleFunc("", h.getAll).Methods("GET")
	legacyRouter.HandleFunc("/withMIP", h.getAllWithMIP).Methods("GET")
}
