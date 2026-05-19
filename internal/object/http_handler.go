package object

import (
	"context"
	"log/slog"
	"net/http"
	"visualizationDbDebet/internal/httpx"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service
}

type service interface {
	getObjectsNameByOrgName(ctx context.Context, orgName string) ([]string, error)
	getObjectByObjectName(ctx context.Context, objectName string) ([]Object, error)
	getObjectsByOrgNameAndObjectName(ctx context.Context, orgName string, objectName string) ([]Object, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getAllObjectsNamesByOrgName(w http.ResponseWriter, r *http.Request) {
	sourceOrgName := mux.Vars(r)["sourceOrgName"]

	names, err := h.service.getObjectsNameByOrgName(r.Context(), sourceOrgName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, names)
	slog.Info("GetAllObjectsNamesByOrgName", "sourceOrgName", sourceOrgName)
}

func (h *Handler) getObjectByName(w http.ResponseWriter, r *http.Request) {
	objectName := r.URL.Query().Get("objectName")

	objects, err := h.service.getObjectByObjectName(r.Context(), objectName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, objects)
	slog.Info("GetObjectByName", "objectName", objectName)
}

func (h *Handler) getAllObjectsByOrgNameAndObjectName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["sourceOrgName"]
	objectName := vars["objectName"]

	objects, err := h.service.getObjectsByOrgNameAndObjectName(r.Context(), sourceOrgName, objectName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, objects)
	slog.Info("GetObjectsByOrgNameAndObjectName", "sourceOrgName", sourceOrgName, "objectName", objectName)
}

func (h *Handler) getAllObjectsByOrgNameAndObjectNameQuery(w http.ResponseWriter, r *http.Request) {
	orgName := r.URL.Query().Get("orgName")
	objectName := r.URL.Query().Get("objectName")

	objects, err := h.service.getObjectsByOrgNameAndObjectName(r.Context(), orgName, objectName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, objects)
	slog.Info("getAllObjectsByOrgNameAndObjectNameQuery", "sourceOrgName", orgName, "objectName", objectName)
}
