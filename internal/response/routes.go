package response

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	responsesRouter := r.PathPrefix("/responses").Subrouter()
	responsesRouter.HandleFunc("", h.GetResponse).Methods("GET")
	responsesRouter.HandleFunc("/with-mip", h.GetResponseWithMIP).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/response").Subrouter()
	legacyRouter.HandleFunc("", h.GetResponse).Methods("GET")
	legacyRouter.HandleFunc("/withMIP", h.GetResponseWithMIP).Methods("GET")
}
