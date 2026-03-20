package util

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, data interface{}) {
	buf, err := json.Marshal(data)
	if err != nil {
		//slog.Error(err.Error(), "info", "Ошибка парса ответа")
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(buf); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError) // ?
		slog.Error("failed to write response", "error", err)
	}
}
