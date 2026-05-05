package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/delivery/util"
	"visualizationBdDebet/internal/object"
)

type ObjectHandler struct {
	service *object.Service
}

func NewObjectHandler(service *object.Service) *ObjectHandler {
	return &ObjectHandler{service: service}
}

func (h *ObjectHandler) GetAllObjectsNamesByOrgName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["sourceOrgName"]

	ctx := r.Context()
	names, err := h.service.GetObjectsNameByOrgName(ctx, sourceOrgName)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("Objects with orgName '%s' not found", sourceOrgName), http.StatusNotFound)
		return
	}

	util.RespondJSON(w, names)
	slog.Info("GetAllObjectsNamesByOrgName", "sourceOrgName", sourceOrgName)
}

func (h *ObjectHandler) GetAllObjectsByOrgNameAndObjectName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["sourceOrgName"]
	objectName := vars["objectName"]

	ctx := r.Context()
	names, err := h.service.GetObjectsByOrgNameAndObjectName(ctx, sourceOrgName, objectName)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("Objects with orgName '%s' or objectName '%s' not found", sourceOrgName, objectName),
			http.StatusNotFound)
		return
	}

	util.RespondJSON(w, names)
	slog.Info("GetObjectsByOrgNameAndObjectName", "sourceOrgName", sourceOrgName, "objectName", objectName)
}

func (h *ObjectHandler) GetAllObjectsByOrgNameAndObjectNameQuery(w http.ResponseWriter, r *http.Request) {
	orgName := r.URL.Query().Get("orgName")
	objectName := r.URL.Query().Get("objectName")

	ctx := r.Context()
	names, err := h.service.GetObjectsByOrgNameAndObjectName(ctx, orgName, objectName)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("Objects with orgName '%s' or objectName '%s' not found", orgName, objectName),
			http.StatusNotFound)
		return
	}

	util.RespondJSON(w, names)
	slog.Info("GetAllObjectsByOrgNameAndObjectNameQuery", "sourceOrgName", orgName, "objectName", objectName)
}
