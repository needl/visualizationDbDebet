package contractor

import (
	"context"
	"errors"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestService_GetContractorsWithBlockFactors_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.GetContractorsWithBlockFactors(ctx, "", "priznanie_bankrotom")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}

	_, err = svc.GetContractorsWithBlockFactors(ctx, "org", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}

func TestService_GetContractorsWithDebt_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.GetContractorsWithDebt(ctx, "", "counterparty")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}

	_, err = svc.GetContractorsWithDebt(ctx, "org", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}

func TestService_GetContractorsWithOverdue_Validation(t *testing.T) {
	svc := NewService(nil)
	ctx := context.Background()

	_, err := svc.GetContractorsWithOverdue(ctx, "", "counterparty")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}

	_, err = svc.GetContractorsWithOverdue(ctx, "org", "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}

func TestService_GetContractorForTable_Validation(t *testing.T) {
	svc := NewService(nil)
	_, err := svc.GetContractorForTable(context.Background(), "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, apperr.ErrInvalidArgument) {
		t.Fatalf("expected ErrInvalidArgument, got %v", err)
	}
}
