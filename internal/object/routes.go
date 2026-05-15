package object

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	objectsRouter := r.PathPrefix("/objects").Subrouter()
	objectsRouter.HandleFunc("/search", h.GetObjectByName).Methods("GET")
	objectsRouter.HandleFunc("/{sourceOrgName}", h.GetAllObjectsNamesByOrgName).Methods("GET")
	objectsRouter.HandleFunc("/{sourceOrgName}/{objectName}", h.GetAllObjectsByOrgNameAndObjectName).Methods("GET")
}
