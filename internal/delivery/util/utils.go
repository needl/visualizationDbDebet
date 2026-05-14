package util

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

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
