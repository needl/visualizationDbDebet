package debet

import (
	"context"
	"errors"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestService_getByOrgName_Validation(t *testing.T) {
	svc := NewService(nil)

	_, err := svc.getByOrgName(context.Background(), "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
	if !strings.Contains(err.Error(), "orgName is required") {
		t.Fatalf("unexpected error message: %v", err)
	}
}
