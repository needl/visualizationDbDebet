package contract

import (
	"context"
	"errors"
	"strings"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestService_GetById_Validation(t *testing.T) {
	t.Parallel()

	svc := NewService(nil)
	tests := []struct {
		name      string
		id        string
		contains  string
		targetErr error
	}{
		{
			name:      "empty id",
			id:        "",
			contains:  "id is required",
			targetErr: apperr.ErrInvalidArgument,
		},
		{
			name:      "non integer id",
			id:        "abc",
			contains:  "id must be integer",
			targetErr: apperr.ErrInvalidArgument,
		},
		{
			name:      "zero id",
			id:        "0",
			contains:  "id must be greater than zero",
			targetErr: apperr.ErrInvalidArgument,
		},
		{
			name:      "negative id",
			id:        "-1",
			contains:  "id must be greater than zero",
			targetErr: apperr.ErrInvalidArgument,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := svc.GetById(context.Background(), tc.id)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !errors.Is(err, tc.targetErr) {
				t.Fatalf("expected %v, got %v", tc.targetErr, err)
			}
			if !strings.Contains(err.Error(), tc.contains) {
				t.Fatalf("expected %q in %q", tc.contains, err.Error())
			}
		})
	}
}
