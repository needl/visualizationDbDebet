package object

import (
	"context"
	"errors"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestService_GetObjectsNameByOrgName_Validation(t *testing.T) {
	svc := NewService(nil)
	_, err := svc.GetObjectsNameByOrgName(context.Background(), "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}

func TestService_GetObjectsByOrgNameAndObjectName_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.GetObjectsByOrgNameAndObjectName(ctx, "", "object")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument for orgName, got %v", err)
	}

	_, err = svc.GetObjectsByOrgNameAndObjectName(ctx, "org", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument for objectName, got %v", err)
	}
}
