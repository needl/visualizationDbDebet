package debet

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	debetsRouter := r.PathPrefix("/debets").Subrouter()
	debetsRouter.HandleFunc("", h.GetAll).Methods("GET")
	debetsRouter.HandleFunc("/with-mip", h.GetAllWithMIP).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/debet").Subrouter()
	legacyRouter.HandleFunc("", h.GetAll).Methods("GET")
	legacyRouter.HandleFunc("/withMIP", h.GetAllWithMIP).Methods("GET")
}
