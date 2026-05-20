package contractor

import (
	"context"
	"errors"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestService_getContractorsWithBlockFactors_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.getContractorsWithBlockFactors(ctx, "", "priznanie_bankrotom")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}

	_, err = svc.getContractorsWithBlockFactors(ctx, "org", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}

func TestService_getContractorsWithDebt_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.getContractorsWithDebt(ctx, "", "counterparty")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}

	_, err = svc.getContractorsWithDebt(ctx, "org", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}

func TestService_getContractorsWithOverdue_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.getContractorsWithOverdue(ctx, "", "counterparty")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}

	_, err = svc.getContractorsWithOverdue(ctx, "org", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}

func TestService_getContractorForTable_Validation(t *testing.T) {
	svc := NewService(nil)
	_, err := svc.getContractorForTable(context.Background(), "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}
