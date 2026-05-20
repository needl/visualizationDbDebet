package response

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	responsesRouter := r.PathPrefix("/responses").Subrouter()
	responsesRouter.HandleFunc("", h.getResponse).Methods("GET")
	responsesRouter.HandleFunc("/with-mip", h.getResponseWithMIP).Methods("GET")

	// Backward-compatible aliases.
	legacyRouter := r.PathPrefix("/response").Subrouter()
	legacyRouter.HandleFunc("", h.getResponse).Methods("GET")
	legacyRouter.HandleFunc("/withMIP", h.getResponseWithMIP).Methods("GET")
}
