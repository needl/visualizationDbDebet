package httpx

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"visualizationDbDebet/internal/apperr"
)

type ErrorMapping struct {
	Err     error
	Status  int
	Message string
}

func RespondJSON(w http.ResponseWriter, data any) {
	buf, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal response", "error", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(buf); err != nil {
		slog.Error("failed to write response", "error", err)
	}
}

func RespondError(w http.ResponseWriter, err error, fallbackMessage string, mappings ...ErrorMapping) {
	if err == nil {
		return
	}

	if errors.Is(err, apperr.ErrInvalidArgument) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, apperr.ErrNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	for _, m := range mappings {
		if m.Err != nil && errors.Is(err, m.Err) {
			http.Error(w, m.Message, m.Status)
			return
		}
	}

	http.Error(w, fallbackMessage, http.StatusInternalServerError)
}
