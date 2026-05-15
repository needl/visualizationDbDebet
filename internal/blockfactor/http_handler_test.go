package blockfactor

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"

	"github.com/gorilla/mux"
)

type blockfactorServiceStub struct {
	getAllFn  func(context.Context) ([]View, error)
	getByIDFn func(context.Context, string) (*View, error)
}

func (s *blockfactorServiceStub) GetAllView(ctx context.Context) ([]View, error) {
	return s.getAllFn(ctx)
}

func (s *blockfactorServiceStub) GetViewById(ctx context.Context, id string) (*View, error) {
	return s.getByIDFn(ctx, id)
}

func TestHandler_GetByID_BadRequest(t *testing.T) {
	h := NewHandler(&blockfactorServiceStub{
		getByIDFn: func(_ context.Context, id string) (*View, error) {
			if id == "" {
				return nil, apperr.NewInvalidArgument("id is required")
			}
			return &View{}, nil
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/blockFactor/", nil)
	rec := httptest.NewRecorder()
	h.GetByID(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "id is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_GetByID_UsesPathVar(t *testing.T) {
	var captured string
	h := NewHandler(&blockfactorServiceStub{
		getByIDFn: func(_ context.Context, id string) (*View, error) {
			captured = id
			return &View{}, nil
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/blockFactor/7", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "7"})
	rec := httptest.NewRecorder()
	h.GetByID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if captured != "7" {
		t.Fatalf("expected id=7, got %q", captured)
	}
}
