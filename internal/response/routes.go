package response

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	responseRouter := r.PathPrefix("/response").Subrouter()
	responseRouter.HandleFunc("", h.GetResponse).Methods("GET")
	responseRouter.HandleFunc("/withMIP", h.GetResponseWithMIP).Methods("GET")
}
