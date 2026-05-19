package blockfactor

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"visualizationDbDebet/internal/apperr"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getAllView(ctx context.Context) ([]View, error) {
	factors, err := s.repo.getViewAll(ctx)
	if err != nil {
		slog.Error("Failed to get all views", "error", err)
		return nil, err
	}

	if len(factors) == 0 {
		slog.Warn("No views found")
		return []View{}, nil
	}

	return factors, nil
}

func (s *Service) getViewByID(ctx context.Context, id string) (*View, error) {
	if id == "" {
		slog.Warn("No view id provided")
		return nil, apperr.NewInvalidArgument("id is required")
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		slog.Warn("Cannot convert blockfactor id")
		return nil, apperr.NewInvalidArgument("id must be integer")
	}

	if intId <= 0 {
		slog.Warn("Invalid view id provided")
		return nil, apperr.NewInvalidArgument("id must be greater than zero")
	}

	factor, err := s.repo.getViewByID(ctx, intId)
	if err != nil {
		slog.Error("Failed to get blockfactor view by id", "error", err)
		return nil, err
	}

	if factor == nil {
		slog.Warn("No view found", "id", id)
		return nil, apperr.NewNotFound(fmt.Sprintf("view with id '%s' not found", id))
	}

	return factor, nil
}
