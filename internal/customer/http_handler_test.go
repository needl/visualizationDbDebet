package customer

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

type customerServiceStub struct{}

func (s *customerServiceStub) getCustomers(context.Context) ([]Customer, error) {
	return []Customer{}, nil
}
func (s *customerServiceStub) getSummaryByCustomerID(_ context.Context, id string) (*Summary, error) {
	if id == "" {
		return nil, apperr.NewInvalidArgument("orgName is required")
	}
	return &Summary{}, nil
}
func (s *customerServiceStub) getTopItemsByCustomerID(_ context.Context, id string) ([]TopItem, error) {
	if id == "" {
		return nil, apperr.NewInvalidArgument("orgName is required")
	}
	return []TopItem{}, nil
}
func (s *customerServiceStub) getTopItemsOverdueByCustomerID(_ context.Context, id string) ([]TopItem, error) {
	if id == "" {
		return nil, apperr.NewInvalidArgument("orgName is required")
	}
	return []TopItem{}, nil
}
func (s *customerServiceStub) getCountBlockFactorsByCustomerID(_ context.Context, id string) (*BlockFactors, error) {
	if id == "" {
		return nil, apperr.NewInvalidArgument("orgName is required")
	}
	return &BlockFactors{}, nil
}

func TestHandler_CustomerEndpoints_RequireOrgNameVar(t *testing.T) {
	h := NewHandler(&customerServiceStub{})

	tests := []struct {
		name   string
		target func(http.ResponseWriter, *http.Request)
	}{
		{name: "summary", target: h.getSummaryByCustomerID},
		{name: "top debtors", target: h.getTopItemsByCustomerID},
		{name: "top overdue", target: h.getTopItemsOverdueByCustomerID},
		{name: "block factors", target: h.getCountBlockFactorsByCustomerID},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/customer", nil)
			rec := httptest.NewRecorder()

			tc.target(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected status 400, got %d", rec.Code)
			}
			if !strings.Contains(rec.Body.String(), "orgName is required") {
				t.Fatalf("unexpected response body: %q", rec.Body.String())
			}
		})
	}
}
