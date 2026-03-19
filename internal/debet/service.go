package debet

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

// GetAllDebets возвращает все записи debet (за исключением "Мосинж")
func (s *Service) GetAllDebets(ctx context.Context) ([]Debet, error) {
	debets, err := s.repo.GetAllDebet(ctx)
	if err != nil {
		slog.Warn("Failed to get debets", "error", err.Error())
		return nil, err
	}

	if len(debets) == 0 {
		slog.Warn("No debets found")
		return []Debet{}, nil
	}

	return debets, nil
}

func (s *Service) GetByOrgName(ctx context.Context, on string) (*Debet, error) {
	debet, err := s.repo.GetByOrgName(ctx, on)
	if err != nil {
		slog.Warn("Failed to get debets", "error", err.Error())
		return nil, err
	}

	if debet == nil {
		slog.Warn("Debet not found")
		return nil, nil
	}

	return debet, nil
}
