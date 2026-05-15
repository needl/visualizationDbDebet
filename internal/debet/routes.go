package debet

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	debetRouter := r.PathPrefix("/debet").Subrouter()
	debetRouter.HandleFunc("", h.GetAll).Methods("GET")
	debetRouter.HandleFunc("/withMIP", h.GetAllWithMIP).Methods("GET")
}
