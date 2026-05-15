package blockfactor

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	blockFactorRouter := r.PathPrefix("/blockFactor").Subrouter()
	blockFactorRouter.HandleFunc("", h.GetAll).Methods("GET")
	blockFactorRouter.HandleFunc("/{id}", h.GetByID).Methods("GET")
}
