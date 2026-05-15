package contract

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"visualizationBdDebet/internal/common"
)

// Service предоставляет бизнес-логику для работы с таблицей contract
type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll(ctx context.Context) ([]View, error) {
	contracts, err := s.repo.GetAllView(ctx)
	if err != nil {
		slog.Warn("Failed to get all contract", "error", err)
		return nil, err
	}

	if len(contracts) == 0 {
		slog.Warn("No contract found")
		return []View{}, nil
	}

	return contracts, nil
}

func (s *Service) GetById(ctx context.Context, id string) (*View, error) {
	if id == "" {
		slog.Warn("Invalid id value")
		return nil, common.NewInvalidArgument("id is required")
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		slog.Warn("Cannot convert contract id to int")
		return nil, common.NewInvalidArgument("id must be integer")
	}

	if intId <= 0 {
		slog.Warn("Invalid id value")
		return nil, common.NewInvalidArgument("id must be greater than zero")
	}

	contract, err := s.repo.GetViewById(ctx, intId)
	if err != nil {
		slog.Warn("Failed to get contract", "error", err)
		return nil, err
	}

	if contract == nil {
		slog.Warn("No contract found", "id", id)
		return nil, common.NewNotFound(fmt.Sprintf("contract with id '%s' not found", id))
	}

	return contract, nil
}
