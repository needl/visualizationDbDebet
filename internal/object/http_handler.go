package object

import (
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/httpx"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllObjectsNamesByOrgName(w http.ResponseWriter, r *http.Request) {
	sourceOrgName := mux.Vars(r)["sourceOrgName"]

	names, err := h.service.GetObjectsNameByOrgName(r.Context(), sourceOrgName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, names)
	slog.Info("GetAllObjectsNamesByOrgName", "sourceOrgName", sourceOrgName)
}

func (h *Handler) GetAllObjectsByOrgNameAndObjectName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceOrgName := vars["sourceOrgName"]
	objectName := vars["objectName"]

	objects, err := h.service.GetObjectsByOrgNameAndObjectName(r.Context(), sourceOrgName, objectName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, objects)
	slog.Info("GetObjectsByOrgNameAndObjectName", "sourceOrgName", sourceOrgName, "objectName", objectName)
}

func (h *Handler) GetAllObjectsByOrgNameAndObjectNameQuery(w http.ResponseWriter, r *http.Request) {
	orgName := r.URL.Query().Get("orgName")
	objectName := r.URL.Query().Get("objectName")

	objects, err := h.service.GetObjectsByOrgNameAndObjectName(r.Context(), orgName, objectName)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, objects)
	slog.Info("GetAllObjectsByOrgNameAndObjectNameQuery", "sourceOrgName", orgName, "objectName", objectName)
}
