package customer

import (
	"context"
	"log/slog"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetCustomers(ctx context.Context) ([]Customer, error) {
	customers, err := s.repo.GetAllCustomers(ctx)
	if err != nil {
		slog.Error("Failed to get customer", "error", err)
		return nil, err
	}

	/*if len(customers) == 0 {
		slog.Warn("No customer")
		return []Customer{}, nil
	}*/

	return customers, nil
}

func (s *Service) GetSummaryByCustomerId(ctx context.Context, id string) (*Summary, error) {
	if id == "" {
		slog.Warn("Customer id is null")
		return nil, nil
	}

	summary, err := s.repo.GetSummaryByCustomerId(ctx, id)
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

	/*if summary == nil {
		slog.Warn("Customer summary is null")
		return nil, nil
	}

	summary.Percentage = summary.TotalAcceptedAmount / summary.TotalContractAmount * 100

	return summary, nil*/
}

func (s *Service) GetTopItemsByCustomerId(ctx context.Context, customerId string) ([]TopItem, error) {
	if customerId == "" {
		slog.Warn("Customer id is null")
		return nil, nil
	}

	items, err := s.repo.GetTopItemsByCustomerId(ctx, customerId)
	if err != nil {
		slog.Warn("Failed to get customer items", "error", err)
		return []TopItem{}, err
	}

	return items, nil
}

func (s *Service) GetTopItemsOverdueByCustomerId(ctx context.Context, customerId string) ([]TopItem, error) {
	if customerId == "" {
		slog.Warn("Customer id is null")
		return nil, nil
	}

	items, err := s.repo.GetTopItemsOverdueByCustomerId(ctx, customerId)
	if err != nil {
		slog.Error("Failed to get customer items", "error", err)
		return []TopItem{}, err
	}

	return items, nil
}

func (s *Service) GetCountBlockFactorsByCustomerId(ctx context.Context, customerId string) (*BlockFactors, error) {
	if customerId == "" {
		slog.Warn("Customer id is null")
		return nil, nil
	}

	factors, err := s.repo.GetCountBlockFactorsByCustomerId(ctx, customerId)
	if err != nil {
		slog.Error("Failed to get customer items", "error", err)
		return nil, err
	}

	return factors, nil
}
