package object

import (
	"context"
	"log/slog"
	"visualizationDbDebet/internal/apperr"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetObjectsNameByOrgName(ctx context.Context, orgName string) ([]string, error) {
	if orgName == "" {
		slog.Warn("orgName is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	names, err := s.repo.FindObjectsNameByOrgName(ctx, orgName)
	if err != nil {
		slog.Error("GetObjectsNameByOrgName err", "err", err)
		return nil, err
	}

	if len(names) == 0 {
		return nil, apperr.NewNotFound("objects not found")
	}

	return names, nil
}

func (s *Service) GetObjectByObjectName(ctx context.Context, objectName string) ([]Object, error) {
	if objectName == "" {
		slog.Warn("objectName is empty")
		return nil, apperr.NewInvalidArgument("objectName is required")
	}

	objects, err := s.repo.FindObjectByName(ctx, objectName)
	if err != nil {
		slog.Error("FindObjectByName err", "err", err)
		return nil, err
	}

	if len(objects) == 0 {
		return nil, apperr.NewNotFound("objects not found")
	}

	return objects, nil
}

func (s *Service) GetObjectsByOrgNameAndObjectName(ctx context.Context, orgName string, objectName string) ([]Object, error) {
	if orgName == "" {
		slog.Warn("orgName is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	if objectName == "" {
		slog.Warn("objectName is empty")
		return nil, apperr.NewInvalidArgument("objectName is required")
	}

	allowedNames, err := s.GetObjectsNameByOrgName(ctx, orgName)
	if err != nil {
		return nil, err
	}

	allowedSet := make(map[string]bool, len(allowedNames))
	for _, name := range allowedNames {
		allowedSet[name] = true
	}

	if !allowedSet[objectName] {
		slog.Warn("Object name not allowed", "name", objectName)
		return nil, apperr.NewNotFound("objects not found")
	}

	objects, err := s.repo.FindObjectsByOrgNameAndObjectName(ctx, orgName, objectName)
	if err != nil {
		slog.Error("FindObjectsByOrgNameAndObjectName err", "err", err)
		return nil, err
	}

	if len(objects) == 0 {
		return nil, apperr.NewNotFound("objects not found")
	}

	return objects, nil
}
