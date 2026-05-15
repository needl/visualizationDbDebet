package apperr

import (
	"errors"
	"strings"
	"testing"
)

func TestNewInvalidArgument(t *testing.T) {
	t.Run("with message", func(t *testing.T) {
		err := NewInvalidArgument("id is required")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrInvalidArgument) {
			t.Fatalf("expected ErrInvalidArgument, got %v", err)
		}
		if !strings.Contains(err.Error(), "id is required") {
			t.Fatalf("expected message in error, got %q", err.Error())
		}
	})

	t.Run("without message", func(t *testing.T) {
		err := NewInvalidArgument("")
		if !errors.Is(err, ErrInvalidArgument) {
			t.Fatalf("expected ErrInvalidArgument, got %v", err)
		}
	})
}

func TestNewNotFound(t *testing.T) {
	t.Run("with message", func(t *testing.T) {
		err := NewNotFound("entity not found")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !errors.Is(err, ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
		if !strings.Contains(err.Error(), "entity not found") {
			t.Fatalf("expected message in error, got %q", err.Error())
		}
	})

	t.Run("without message", func(t *testing.T) {
		err := NewNotFound("")
		if !errors.Is(err, ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
}
