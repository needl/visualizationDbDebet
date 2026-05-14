package util

import (
	"errors"
	"net/http"
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

	for _, m := range mappings {
		if m.Err != nil && errors.Is(err, m.Err) {
			http.Error(w, m.Message, m.Status)
			return
		}
	}

	http.Error(w, fallbackMessage, http.StatusInternalServerError)
}
