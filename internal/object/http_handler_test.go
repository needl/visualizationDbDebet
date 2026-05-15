package object

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

type objectServiceStub struct{}

func (s *objectServiceStub) GetObjectsNameByOrgName(_ context.Context, orgName string) ([]string, error) {
	if orgName == "" {
		return nil, apperr.NewInvalidArgument("orgName is required")
	}
	return []string{"obj"}, nil
}

func (s *objectServiceStub) GetObjectByObjectName(_ context.Context, objectName string) ([]Object, error) {
	if objectName == "" {
		return nil, apperr.NewInvalidArgument("objectName is required")
	}
	return []Object{}, nil
}

func (s *objectServiceStub) GetObjectsByOrgNameAndObjectName(_ context.Context, orgName string, objectName string) ([]Object, error) {
	if orgName == "" {
		return nil, apperr.NewInvalidArgument("orgName is required")
	}
	if objectName == "" {
		return nil, apperr.NewInvalidArgument("objectName is required")
	}
	return []Object{}, nil
}

func TestHandler_GetAllObjectsNamesByOrgName_BadRequest(t *testing.T) {
	h := NewHandler(&objectServiceStub{})

	req := httptest.NewRequest(http.MethodGet, "/objects/", nil)
	rec := httptest.NewRecorder()
	h.GetAllObjectsNamesByOrgName(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "orgName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_GetAllObjectsByOrgNameAndObjectName_BadRequest(t *testing.T) {
	h := NewHandler(&objectServiceStub{})

	req := httptest.NewRequest(http.MethodGet, "/objects/org/", nil)
	rec := httptest.NewRecorder()
	h.GetAllObjectsByOrgNameAndObjectName(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "orgName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_GetAllObjectsByOrgNameAndObjectNameQuery_BadRequest(t *testing.T) {
	h := NewHandler(&objectServiceStub{})

	req := httptest.NewRequest(http.MethodGet, "/objects/search", nil)
	rec := httptest.NewRecorder()
	h.GetAllObjectsByOrgNameAndObjectNameQuery(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "orgName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}
