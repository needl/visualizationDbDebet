package debet

import (
	"context"
	"log/slog"
)

// Service предоставляет бизнес-логику для работы с таблицей debet
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
		slog.Warn("Failed to get debets", "error", err)
		return nil, err
	}

	if len(debets) == 0 {
		slog.Warn("No debets found")
		return []Debet{}, nil
	}

	return debets, nil
}

// GetByOrgName возвращается запись из таблицы debet по названию организации
func (s *Service) GetByOrgName(ctx context.Context, orgName string) (*Debet, error) {
	if orgName == "" {
		slog.Warn("OrgName is empty")
		return nil, nil
	}

	debet, err := s.repo.GetByOrgName(ctx, orgName)
	if err != nil {
		slog.Warn("Failed to get debet by orgName", "error", err)
		return nil, err
	}

	if debet == nil {
		slog.Warn("Debet not found")
		return nil, nil
	}

	return debet, nil
}
