package response

import (
	"context"
	"fmt"
	"visualizationBdDebet/internal/debet"
)

type Service struct {
	repo         *Repository
	debetService *debet.Service
}

func NewService(repo *Repository, debetService *debet.Service) *Service {
	return &Service{repo: repo, debetService: debetService}
}

// GetResponse получает сводную статистику по дебиторке без МИП
func (s *Service) GetResponse(ctx context.Context) (*Response, error) {
	stats, err := s.repo.GetPageDto(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get response stats from repo in response service: %w", err)
	}

	return &Response{
		CountSourceOrg:    stats.CountSourceOrg,
		CountContracts:    stats.CountContracts,
		SumContractAmount: stats.SumContractAmount,
		SumDebetTotal:     stats.SumDebetTotal,
		SumDebetOverdue:   stats.SumDebetOverdue,
	}, nil
}

// GetResponseWithMIP получает сводную статистику по дебиторке с МИП
func (s *Service) GetResponseWithMIP(ctx context.Context) (*Response, error) {
	stats, err := s.repo.GetPageDtoWithMIP(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get response stats from repo in response service: %w", err)
	}

	return &Response{
		CountSourceOrg:    stats.CountSourceOrg,
		CountContracts:    stats.CountContracts,
		SumContractAmount: stats.SumContractAmount,
		SumDebetTotal:     stats.SumDebetTotal,
		SumDebetOverdue:   stats.SumDebetOverdue,
	}, nil
}
