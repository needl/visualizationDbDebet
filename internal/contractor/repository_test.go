package contractor

import (
	"context"
	"errors"
	"testing"
)

func TestRepository_FindContractorsWithBlockFactors_InvalidColumn(t *testing.T) {
	repo := NewRepository(nil)

	_, err := repo.FindContractorsWithBlockFactors(context.Background(), "org", "unknown_column")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrColumnNotAllowed) {
		t.Fatalf("expected ErrColumnNotAllowed, got %v", err)
	}
}
