package httpx

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestRespondJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		RespondJSON(rec, map[string]string{"status": "ok"})

		if rec.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", rec.Code)
		}
		if got := rec.Header().Get("Content-Type"); got != "application/json" {
			t.Fatalf("expected Content-Type application/json, got %q", got)
		}

		var payload map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}
		if payload["status"] != "ok" {
			t.Fatalf("expected status=ok, got %q", payload["status"])
		}
	})

	t.Run("marshal error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		RespondJSON(rec, func() {})

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("expected status 500, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "failed to encode response") {
			t.Fatalf("unexpected body: %q", rec.Body.String())
		}
	})
}

func TestRespondError(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		RespondError(rec, nil, "fallback")
		if rec.Body.Len() != 0 {
			t.Fatalf("expected empty body, got %q", rec.Body.String())
		}
	})

	t.Run("invalid argument", func(t *testing.T) {
		rec := httptest.NewRecorder()
		err := apperr.NewInvalidArgument("id is required")
		RespondError(rec, err, "fallback")

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "id is required") {
			t.Fatalf("unexpected body: %q", rec.Body.String())
		}
	})

	t.Run("not found", func(t *testing.T) {
		rec := httptest.NewRecorder()
		err := apperr.NewNotFound("not found")
		RespondError(rec, err, "fallback")

		if rec.Code != http.StatusNotFound {
			t.Fatalf("expected status 404, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "not found") {
			t.Fatalf("unexpected body: %q", rec.Body.String())
		}
	})

	t.Run("custom mapping", func(t *testing.T) {
		rec := httptest.NewRecorder()
		errConflict := errors.New("conflict")
		err := errors.Join(errors.New("wrap"), errConflict)
		RespondError(rec, err, "fallback", ErrorMapping{
			Err:     errConflict,
			Status:  http.StatusConflict,
			Message: "custom conflict",
		})

		if rec.Code != http.StatusConflict {
			t.Fatalf("expected status 409, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "custom conflict") {
			t.Fatalf("unexpected body: %q", rec.Body.String())
		}
	})

	t.Run("fallback", func(t *testing.T) {
		rec := httptest.NewRecorder()
		RespondError(rec, errors.New("boom"), "internal server error")

		if rec.Code != http.StatusInternalServerError {
			t.Fatalf("expected status 500, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "internal server error") {
			t.Fatalf("unexpected body: %q", rec.Body.String())
		}
	})
}
