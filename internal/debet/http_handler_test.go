package debet

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"

	"github.com/gorilla/mux"
)

type debetServiceStub struct {
	getAllFn        func(context.Context) ([]View, error)
	getAllWithMIPFn func(context.Context) ([]View, error)
	getByOrgNameFn  func(context.Context, string) (*View, error)
}

func (s *debetServiceStub) getAll(ctx context.Context) ([]View, error) {
	return s.getAllFn(ctx)
}

func (s *debetServiceStub) getAllWithMIP(ctx context.Context) ([]View, error) {
	return s.getAllWithMIPFn(ctx)
}

func (s *debetServiceStub) getByOrgName(ctx context.Context, orgName string) (*View, error) {
	return s.getByOrgNameFn(ctx, orgName)
}

func TestHandler_getByOrgName_BadRequest(t *testing.T) {
	h := NewHandler(&debetServiceStub{
		getByOrgNameFn: func(_ context.Context, orgName string) (*View, error) {
			if orgName == "" {
				return nil, apperr.NewInvalidArgument("orgName is required")
			}
			return &View{}, nil
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/debet/", nil)
	rec := httptest.NewRecorder()
	h.getByOrgName(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "orgName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_getByOrgName_UsesPathVar(t *testing.T) {
	var captured string
	h := NewHandler(&debetServiceStub{
		getByOrgNameFn: func(_ context.Context, orgName string) (*View, error) {
			captured = orgName
			return &View{OrgName: orgName}, nil
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/debet/acme", nil)
	req = mux.SetURLVars(req, map[string]string{"orgName": "acme"})
	rec := httptest.NewRecorder()
	h.getByOrgName(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if captured != "acme" {
		t.Fatalf("expected orgName acme, got %q", captured)
	}
}
