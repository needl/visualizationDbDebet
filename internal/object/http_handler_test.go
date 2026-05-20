package object

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

type objectServiceStub struct {
	calledByName       bool
	calledByOrgAndName bool
	calledByCtrAndName bool
}

func (s *objectServiceStub) getObjectsNameByOrgName(_ context.Context, orgName string) ([]string, error) {
	if orgName == "" {
		return nil, apperr.NewInvalidArgument("orgName is required")
	}
	return []string{"obj"}, nil
}

func (s *objectServiceStub) getObjectByObjectName(_ context.Context, objectName string) ([]Object, error) {
	s.calledByName = true

	if objectName == "" {
		return nil, apperr.NewInvalidArgument("objectName is required")
	}
	return []Object{}, nil
}

func (s *objectServiceStub) getObjectsByOrgNameAndObjectName(_ context.Context, orgName string, objectName string) ([]Object, error) {
	s.calledByOrgAndName = true

	if orgName == "" {
		return nil, apperr.NewInvalidArgument("orgName is required")
	}
	if objectName == "" {
		return nil, apperr.NewInvalidArgument("objectName is required")
	}
	return []Object{}, nil
}

func (s *objectServiceStub) getObjectsByCounterpartyNameAndObjectName(_ context.Context, counterpartyName string, objectName string) ([]Object, error) {
	s.calledByCtrAndName = true

	if counterpartyName == "" {
		return nil, apperr.NewInvalidArgument("counterpartyName is required")
	}
	if objectName == "" {
		return nil, apperr.NewInvalidArgument("objectName is required")
	}
	return []Object{}, nil
}

func TestHandler_getAllObjectsNamesByOrgName_BadRequest(t *testing.T) {
	h := NewHandler(&objectServiceStub{})

	req := httptest.NewRequest(http.MethodGet, "/objects/", nil)
	rec := httptest.NewRecorder()
	h.getAllObjectsNamesByOrgName(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "orgName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_getObjectByName_WithCounterpartyName_UsesScopedSearch(t *testing.T) {
	stub := &objectServiceStub{}
	h := NewHandler(stub)

	req := httptest.NewRequest(http.MethodGet, "/objects/search?counterpartyName=ctr&objectName=obj", nil)
	rec := httptest.NewRecorder()
	h.getObjectByName(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !stub.calledByCtrAndName {
		t.Fatalf("expected scoped lookup by counterpartyName and objectName")
	}
	if stub.calledByName || stub.calledByOrgAndName {
		t.Fatalf("unexpected lookup with wrong method")
	}
}

func TestHandler_getObjectByName_WithoutCounterpartyName_ReturnsBadRequest(t *testing.T) {
	stub := &objectServiceStub{}
	h := NewHandler(stub)

	req := httptest.NewRequest(http.MethodGet, "/objects/search?objectName=obj", nil)
	rec := httptest.NewRecorder()
	h.getObjectByName(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "counterpartyName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
	if stub.calledByName || stub.calledByOrgAndName || stub.calledByCtrAndName {
		t.Fatalf("service should not be called when counterpartyName is missing")
	}
}

func TestHandler_getAllObjectsByOrgNameAndObjectName_BadRequest(t *testing.T) {
	h := NewHandler(&objectServiceStub{})

	req := httptest.NewRequest(http.MethodGet, "/objects/org/", nil)
	rec := httptest.NewRecorder()
	h.getAllObjectsByOrgNameAndObjectName(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "orgName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_getAllObjectsByOrgNameAndObjectNameQuery_BadRequest(t *testing.T) {
	h := NewHandler(&objectServiceStub{})

	req := httptest.NewRequest(http.MethodGet, "/objects/search", nil)
	rec := httptest.NewRecorder()
	h.getAllObjectsByOrgNameAndObjectNameQuery(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "orgName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}
