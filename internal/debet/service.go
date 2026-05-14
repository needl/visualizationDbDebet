package debet

import (
	"context"
	"log/slog"
	"visualizationBdDebet/internal/common"
)

// Service предоставляет бизнес-логику для работы с таблицей debet
type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetAll возвращает все записи View debet (за исключением "Мосинж")
func (s *Service) GetAll(ctx context.Context) ([]View, error) {
	debets, err := s.repo.GetAllView(ctx)
	if err != nil {
		slog.Error("Failed to get debets", "error", err)
		return nil, err
	}

	if len(debets) == 0 {
		slog.Warn("No debets found")
		return []View{}, nil
	}

	return debets, nil
}

// GetAllWithMIP возвращает все записи View debet
func (s *Service) GetAllWithMIP(ctx context.Context) ([]View, error) {
	debets, err := s.repo.GetAllViewWithMIP(ctx)
	if err != nil {
		slog.Error("Failed to get debets", "error", err)
		return nil, err
	}

	if len(debets) == 0 {
		slog.Warn("No debets found")
		return []View{}, nil
	}

	return debets, nil
}

// GetByOrgName возвращается запись из таблицы debet по названию организации
func (s *Service) GetByOrgName(ctx context.Context, orgName string) (*View, error) {
	if orgName == "" {
		slog.Warn("OrgName is empty")
		return nil, common.NewInvalidArgument("orgName is required")
	}

	debet, err := s.repo.GetViewByOrgName(ctx, orgName)
	if err != nil {
		slog.Warn("Failed to get debet by orgName", "error", err)
		return nil, err
	}

	if debet == nil {
		slog.Warn("View not found", "orgName", orgName)
		return nil, common.NewNotFound("debet not found")
	}

	return debet, nil
}
