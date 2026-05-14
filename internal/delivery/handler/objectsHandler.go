package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/delivery/util"
	"visualizationBdDebet/internal/object"

	"github.com/gorilla/mux"
)

type ObjectHandler struct {
	service *object.Service
}

func NewHandlerObject(service *object.Service) *ObjectHandler {
	return &ObjectHandler{service: service}
}

func (h *ObjectHandler) GetAllObjectsNamesByOrgName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["sourceOrgName"]

	ctx := r.Context()
	names, err := h.service.GetObjectsNameByOrgName(ctx, sourceOrgName)
	if err != nil {
		util.RespondError(
			w,
			err,
			"internal server error",
			util.ErrorMapping{
				Err:     object.ErrOrgNameEmpty,
				Status:  http.StatusBadRequest,
				Message: "sourceOrgName is required",
			},
			util.ErrorMapping{
				Err:     object.ErrObjectsNotFound,
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("Objects with orgName '%s' not found", sourceOrgName),
			},
		)
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
		util.RespondError(
			w,
			err,
			"internal server error",
			util.ErrorMapping{
				Err:     object.ErrOrgNameEmpty,
				Status:  http.StatusBadRequest,
				Message: "orgName and objectName are required",
			},
			util.ErrorMapping{
				Err:     object.ErrObjectNameEmpty,
				Status:  http.StatusBadRequest,
				Message: "orgName and objectName are required",
			},
			util.ErrorMapping{
				Err:     object.ErrObjectNameNotAllowed,
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("Objects with orgName '%s' or objectName '%s' not found", sourceOrgName, objectName),
			},
			util.ErrorMapping{
				Err:     object.ErrObjectsNotFound,
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("Objects with orgName '%s' or objectName '%s' not found", sourceOrgName, objectName),
			},
		)
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
		util.RespondError(
			w,
			err,
			"internal server error",
			util.ErrorMapping{
				Err:     object.ErrOrgNameEmpty,
				Status:  http.StatusBadRequest,
				Message: "orgName and objectName are required",
			},
			util.ErrorMapping{
				Err:     object.ErrObjectNameEmpty,
				Status:  http.StatusBadRequest,
				Message: "orgName and objectName are required",
			},
			util.ErrorMapping{
				Err:     object.ErrObjectNameNotAllowed,
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("Objects with orgName '%s' or objectName '%s' not found", orgName, objectName),
			},
			util.ErrorMapping{
				Err:     object.ErrObjectsNotFound,
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("Objects with orgName '%s' or objectName '%s' not found", orgName, objectName),
			},
		)
		return
	}

	util.RespondJSON(w, names)
	slog.Info("GetAllObjectsByOrgNameAndObjectNameQuery", "sourceOrgName", orgName, "objectName", objectName)
}
