package customer

import (
	"context"
	"errors"
	"testing"
	"visualizationDbDebet/internal/apperr"
)

func TestService_validationEmptyCustomerID(t *testing.T) {
	ctx := context.Background()
	svc := NewService(nil)

	cases := []struct {
		name string
		call func() error
	}{
		{
			name: "getSummaryByCustomerID",
			call: func() error {
				_, err := svc.getSummaryByCustomerID(ctx, "")
				return err
			},
		},
		{
			name: "getTopItemsByCustomerID",
			call: func() error {
				_, err := svc.getTopItemsByCustomerID(ctx, "")
				return err
			},
		},
		{
			name: "getTopItemsOverdueByCustomerID",
			call: func() error {
				_, err := svc.getTopItemsOverdueByCustomerID(ctx, "")
				return err
			},
		},
		{
			name: "getCountBlockFactorsByCustomerID",
			call: func() error {
				_, err := svc.getCountBlockFactorsByCustomerID(ctx, "")
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
