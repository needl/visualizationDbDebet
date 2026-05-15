package response

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type responseServiceStub struct {
	getResponseFn        func(context.Context) (*Response, error)
	getResponseWithMIPFn func(context.Context) (*Response, error)
}

func (s *responseServiceStub) GetResponse(ctx context.Context) (*Response, error) {
	return s.getResponseFn(ctx)
}

func (s *responseServiceStub) GetResponseWithMIP(ctx context.Context) (*Response, error) {
	return s.getResponseWithMIPFn(ctx)
}

func TestHandler_GetResponse_Success(t *testing.T) {
	h := NewHandler(&responseServiceStub{
		getResponseFn: func(context.Context) (*Response, error) {
			return &Response{CountSourceOrg: 2}, nil
		},
		getResponseWithMIPFn: func(context.Context) (*Response, error) { return &Response{}, nil },
	})

	req := httptest.NewRequest(http.MethodGet, "/response", nil)
	rec := httptest.NewRecorder()
	h.GetResponse(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"count_source_org":2`) {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}

func TestHandler_GetResponseWithMIP_InternalError(t *testing.T) {
	h := NewHandler(&responseServiceStub{
		getResponseFn: func(context.Context) (*Response, error) { return &Response{}, nil },
		getResponseWithMIPFn: func(context.Context) (*Response, error) {
			return nil, errors.New("db down")
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/response/withMIP", nil)
	rec := httptest.NewRecorder()
	h.GetResponseWithMIP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "internal server error") {
		t.Fatalf("unexpected response body: %q", rec.Body.String())
	}
}
