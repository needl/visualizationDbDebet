package customer

import (
	"context"
	"log/slog"
	"visualizationDbDebet/internal/apperr"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getCustomers(ctx context.Context) ([]Customer, error) {
	customers, err := s.repo.findAllCustomers(ctx)
	if err != nil {
		slog.Error("Failed to get customer", "error", err)
		return nil, err
	}

	return customers, nil
}

func (s *Service) getSummaryByCustomerID(ctx context.Context, id string) (*Summary, error) {
	if id == "" {
		slog.Warn("Customer id is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	summary, err := s.repo.findSummaryByCustomerID(ctx, id)
	if err != nil {
		slog.Error("Failed to get customer summary", "error", err)
		return nil, err
	}

	if summary.TotalContractAmount != 0 && summary.TotalAcceptedAmount != 0 {
		summary.Percentage = summary.TotalAcceptedAmount / summary.TotalContractAmount * 100
	} else {
		summary.Percentage = 0
	}
	return summary, nil
}

func (s *Service) getTopItemsByCustomerID(ctx context.Context, customerID string) ([]TopItem, error) {
	if customerID == "" {
		slog.Warn("Customer id is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	items, err := s.repo.findTopItemsByCustomerID(ctx, customerID)
	if err != nil {
		slog.Warn("Failed to get customer items", "error", err)
		return []TopItem{}, err
	}

	return items, nil
}

func (s *Service) getTopItemsOverdueByCustomerID(ctx context.Context, customerID string) ([]TopItem, error) {
	if customerID == "" {
		slog.Warn("Customer id is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	items, err := s.repo.findTopItemsOverdueByCustomerID(ctx, customerID)
	if err != nil {
		slog.Error("Failed to get customer items", "error", err)
		return []TopItem{}, err
	}

	return items, nil
}

func (s *Service) getCountBlockFactorsByCustomerID(ctx context.Context, customerID string) (*BlockFactors, error) {
	if customerID == "" {
		slog.Warn("Customer id is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	factors, err := s.repo.findCountBlockFactorsByCustomerID(ctx, customerID)
	if err != nil {
		slog.Error("Failed to get customer items", "error", err)
		return nil, err
	}

	return factors, nil
}
