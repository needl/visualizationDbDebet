package contractor

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandler_GetContractorsWithBlockFactors_BadRequest(t *testing.T) {
	h := NewHandler(NewService(nil))

	t.Run("missing orgName from path vars", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/contractor/?columnName=priznanie_bankrotom", nil)
		rec := httptest.NewRecorder()

		h.GetContractorsWithBlockFactors(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "orgName is required") {
			t.Fatalf("unexpected response body: %q", rec.Body.String())
		}
	})

	t.Run("missing columnName query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/contractor/org", nil)
		req = mux.SetURLVars(req, map[string]string{"orgName": "org"})
		rec := httptest.NewRecorder()

		h.GetContractorsWithBlockFactors(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "columnName is required") {
			t.Fatalf("unexpected response body: %q", rec.Body.String())
		}
	})
}

func TestHandler_GetContractorsWithDebt_BadRequest(t *testing.T) {
	h := NewHandler(NewService(nil))

	req := httptest.NewRequest(http.MethodGet, "/contractor/org/debt", nil)
	req = mux.SetURLVars(req, map[string]string{"orgName": "org"})
	rec := httptest.NewRecorder()

	h.GetContractorsWithDebt(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "counterpartyName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_GetContractorsWithOverdue_BadRequest(t *testing.T) {
	h := NewHandler(NewService(nil))

	req := httptest.NewRequest(http.MethodGet, "/contractor/org/overdue", nil)
	req = mux.SetURLVars(req, map[string]string{"orgName": "org"})
	rec := httptest.NewRecorder()

	h.GetContractorsWithOverdue(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "counterpartyName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_GetContractorForTable_BadRequest(t *testing.T) {
	h := NewHandler(NewService(nil))

	req := httptest.NewRequest(http.MethodGet, "/contractor/table", nil)
	rec := httptest.NewRecorder()

	h.GetContractorForTable(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "counterpartyName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}
