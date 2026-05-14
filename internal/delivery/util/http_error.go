package util

import (
	"errors"
	"net/http"
	"visualizationBdDebet/internal/common"
)

type ErrorMapping struct {
	Err     error
	Status  int
	Message string
}

func RespondError(w http.ResponseWriter, err error, fallbackMessage string, mappings ...ErrorMapping) {
	if err == nil {
		return
	}

	if errors.Is(err, common.ErrInvalidArgument) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, common.ErrNotFound) {
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
