package customer

import (
	"context"
	"errors"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestService_Validation_EmptyCustomerID(t *testing.T) {
	ctx := context.Background()
	svc := NewService(nil)

	cases := []struct {
		name string
		call func() error
	}{
		{
			name: "GetSummaryByCustomerId",
			call: func() error {
				_, err := svc.GetSummaryByCustomerId(ctx, "")
				return err
			},
		},
		{
			name: "GetTopItemsByCustomerId",
			call: func() error {
				_, err := svc.GetTopItemsByCustomerId(ctx, "")
				return err
			},
		},
		{
			name: "GetTopItemsOverdueByCustomerId",
			call: func() error {
				_, err := svc.GetTopItemsOverdueByCustomerId(ctx, "")
				return err
			},
		},
		{
			name: "GetCountBlockFactorsByCustomerId",
			call: func() error {
				_, err := svc.GetCountBlockFactorsByCustomerId(ctx, "")
				return err
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.call()
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !errors.Is(err, apperr.ErrInvalidArgument) {
				t.Fatalf("expected ErrInvalidArgument, got %v", err)
			}
		})
	}
}
