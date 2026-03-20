package blockfactor

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllView(ctx context.Context) ([]View, error) {
	factors, err := s.repo.GetViewAll(ctx)
	if err != nil {
		slog.Error("Failed to get all views", "error", err)
		return nil, err
	}

	if len(factors) == 0 {
		slog.Error("No views found")
		return []View{}, nil
	}

	return factors, nil
}

func (s *Service) GetViewById(ctx context.Context, id string) (*View, error) {
	if id == "" {
		slog.Error("No view id provided")
		return nil, nil
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		slog.Error("Cant convert blockfactor id provided")
		return nil, err
	}

	if intId <= 0 {
		slog.Error("No view id provided")
		return nil, errors.New("invalid view id")
	}

	factor, err := s.repo.GetViewById(ctx, intId)
	if err != nil {
		slog.Error("Failed to get all views: %v", "error", err)
		return nil, err
	}

	if factor == nil {
		slog.Warn("No view found", "id", id)
		return nil, nil
	}

	return factor, nil
}
