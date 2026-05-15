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

func (s *Service) GetCustomers(ctx context.Context) ([]Customer, error) {
	customers, err := s.repo.FindAllCustomers(ctx)
	if err != nil {
		slog.Error("Failed to get customer", "error", err)
		return nil, err
	}

	return customers, nil
}

func (s *Service) GetSummaryByCustomerId(ctx context.Context, id string) (*Summary, error) {
	if id == "" {
		slog.Warn("Customer id is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	summary, err := s.repo.FindSummaryByCustomerId(ctx, id)
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

func (s *Service) GetTopItemsByCustomerId(ctx context.Context, customerId string) ([]TopItem, error) {
	if customerId == "" {
		slog.Warn("Customer id is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	items, err := s.repo.FindTopItemsByCustomerId(ctx, customerId)
	if err != nil {
		slog.Warn("Failed to get customer items", "error", err)
		return []TopItem{}, err
	}

	return items, nil
}

func (s *Service) GetTopItemsOverdueByCustomerId(ctx context.Context, customerId string) ([]TopItem, error) {
	if customerId == "" {
		slog.Warn("Customer id is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	items, err := s.repo.FindTopItemsOverdueByCustomerId(ctx, customerId)
	if err != nil {
		slog.Error("Failed to get customer items", "error", err)
		return []TopItem{}, err
	}

	return items, nil
}

func (s *Service) GetCountBlockFactorsByCustomerId(ctx context.Context, customerId string) (*BlockFactors, error) {
	if customerId == "" {
		slog.Warn("Customer id is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	factors, err := s.repo.FindCountBlockFactorsByCustomerId(ctx, customerId)
	if err != nil {
		slog.Error("Failed to get customer items", "error", err)
		return nil, err
	}

	return factors, nil
}
