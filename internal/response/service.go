package response

import (
	"context"
	"fmt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetResponse получает сводную статистику по дебиторке без МИП
func (s *Service) GetResponse(ctx context.Context) (*Response, error) {
	stats, err := s.repo.GetPageDto(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get response stats from repo in response service: %w", err)
	}

	return stats, nil
}

// GetResponseWithMIP получает сводную статистику по дебиторке с МИП
func (s *Service) GetResponseWithMIP(ctx context.Context) (*Response, error) {
	stats, err := s.repo.GetPageDtoWithMIP(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get response stats from repo in response service: %w", err)
	}

	return stats, nil
}
