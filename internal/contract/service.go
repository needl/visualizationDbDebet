package contract

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
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
		return nil, errors.New("invalid id value")
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		slog.Warn("Cant convert id to int")
		return nil, err
	}

	if intId <= 0 {
		slog.Warn("Invalid id value")
		return nil, errors.New("invalid id value")
	}

	contract, err := s.repo.GetViewById(ctx, intId)
	if err != nil {
		slog.Warn("Failed to get contract", "error", err)
		return nil, err
	}

	if contract == nil {
		slog.Warn("No contract found", "id", id)
		return nil, nil
	}

	return contract, nil
}
