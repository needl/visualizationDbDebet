package contractoranalysis

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"

	"github.com/gorilla/mux"
)

type contractorAnalysisServiceStub struct {
	analyticsErr error
}

func (s *contractorAnalysisServiceStub) getContractors(context.Context) ([]string, error) {
	return []string{"ООО Пума"}, nil
}

func (s *contractorAnalysisServiceStub) getAnalytics(_ context.Context, contractorName string) (*Analytics, error) {
	if s.analyticsErr != nil {
		return nil, s.analyticsErr
	}

	if contractorName == "" {
		return nil, apperr.NewInvalidArgument("contractorName is required")
	}

	return &Analytics{
		ContractorName: contractorName,
		Summary: Summary{
			ContractsSum:        1,
			ObjectsCount:        1,
			OverdueObjectsCount: 0,
		},
	}, nil
}

func (s *contractorAnalysisServiceStub) getObjectDetails(
	_ context.Context,
	contractorName string,
	customerName string,
	objectName string,
) (*ObjectDetails, error) {
	if contractorName == "" {
		return nil, apperr.NewInvalidArgument("contractorName is required")
	}
	if customerName == "" {
		return nil, apperr.NewInvalidArgument("customerName is required")
	}
	if objectName == "" {
		return nil, apperr.NewInvalidArgument("objectName is required")
	}

	return &ObjectDetails{
		CustomerName:      customerName,
		ContractorName:    contractorName,
		ObjectName:        objectName,
		OverdueDebtAmount: 0,
	}, nil
}

func TestContractorAnalysisRoutes_GetContractors(t *testing.T) {
	h := NewHandler(&contractorAnalysisServiceStub{})
	r := mux.NewRouter()
	RegisterRoutes(r, h)

	req := httptest.NewRequest(http.MethodGet, "/contractor-analytics", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "ООО Пума") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestContractorAnalysisRoutes_GetAnalytics(t *testing.T) {
	h := NewHandler(&contractorAnalysisServiceStub{})
	r := mux.NewRouter()
	RegisterRoutes(r, h)

	req := httptest.NewRequest(http.MethodGet, "/contractor-analytics/ООО%20Пума", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "\"contractor_name\":\"ООО Пума\"") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestContractorAnalysisRoutes_GetAnalytics_NoObjects(t *testing.T) {
	h := NewHandler(&contractorAnalysisServiceStub{
		analyticsErr: fmt.Errorf("%w: АО \"МОСЭНЕРГОСБЫТ\"", errContractorHasNoObjects),
	})
	r := mux.NewRouter()
	RegisterRoutes(r, h)

	req := httptest.NewRequest(http.MethodGet, "/contractor-analytics/АО%20%22МОСЭНЕРГОСБЫТ%22", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected status %d, got %d", http.StatusUnprocessableEntity, rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "contractor has no objects") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestContractorAnalysisRoutes_GetObjectDetails(t *testing.T) {
	h := NewHandler(&contractorAnalysisServiceStub{})
	r := mux.NewRouter()
	RegisterRoutes(r, h)

	req := httptest.NewRequest(http.MethodGet, "/contractor-analytics/ООО%20Пума/object-details?customerName=Заказчик%201&objectName=ЖК%20Северный", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "\"customer_name\":\"Заказчик 1\"") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestContractorAnalysisRoutes_GetObjectDetails_BadRequest(t *testing.T) {
	h := NewHandler(&contractorAnalysisServiceStub{})
	r := mux.NewRouter()
	RegisterRoutes(r, h)

	req := httptest.NewRequest(http.MethodGet, "/contractor-analytics/ООО%20Пума/object-details?objectName=ЖК%20Северный", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "customerName is required") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}
