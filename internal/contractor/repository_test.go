package contractor

import (
	"context"
	"errors"
	"testing"
)

func TestRepository_findContractorsWithBlockFactors_InvalidColumn(t *testing.T) {
	repo := NewRepository(nil)

	_, err := repo.findContractorsWithBlockFactors(context.Background(), "org", "unknown_column")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, errColumnNotAllowed) {
		t.Fatalf("expected errColumnNotAllowed, got %v", err)
	}
}
