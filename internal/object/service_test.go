package object

import (
	"context"
	"errors"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestService_getObjectsNameByOrgName_Validation(t *testing.T) {
	svc := NewService(nil)
	_, err := svc.getObjectsNameByOrgName(context.Background(), "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}

func TestService_getObjectsByOrgNameAndObjectName_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.getObjectsByOrgNameAndObjectName(ctx, "", "object")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument for orgName, got %v", err)
	}

	_, err = svc.getObjectsByOrgNameAndObjectName(ctx, "org", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument for objectName, got %v", err)
	}
}

func TestService_getObjectsByCounterpartyNameAndObjectName_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.getObjectsByCounterpartyNameAndObjectName(ctx, "", "object")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument for counterpartyName, got %v", err)
	}

	_, err = svc.getObjectsByCounterpartyNameAndObjectName(ctx, "counterparty", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument for objectName, got %v", err)
	}
}
