package object

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	objectsRouter := r.PathPrefix("/objects").Subrouter()
	objectsRouter.HandleFunc("/search", h.getObjectByName).Methods("GET")
	objectsRouter.HandleFunc("/{sourceOrgName}", h.getAllObjectsNamesByOrgName).Methods("GET")
	objectsRouter.HandleFunc("/{sourceOrgName}/{objectName}", h.getAllObjectsByOrgNameAndObjectName).Methods("GET")
}
