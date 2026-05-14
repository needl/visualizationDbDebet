package object

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
)

type Service struct {
	repo *Repository
}

var (
	ErrOrgNameEmpty         = errors.New("orgName is empty")
	ErrObjectNameEmpty      = errors.New("objectName is empty")
	ErrObjectNameNotAllowed = errors.New("object name not allowed")
	ErrObjectsNotFound      = errors.New("no objects")
)

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetObjectsNameByOrgName(ctx context.Context, orgName string) ([]string, error) {
	if orgName == "" {
		slog.Warn("orgName is empty")
		return nil, ErrOrgNameEmpty
	}

	names, err := s.repo.FindObjectsNameByOrgName(ctx, orgName)
	if err != nil {
		slog.Error("GetObjectsNameByOrgName err", "err", err)
		return nil, err
	}

	if len(names) == 0 {
		return nil, ErrObjectsNotFound
	}

	return names, nil
}

func (s *Service) GetObjectsByOrgNameAndObjectName(ctx context.Context, orgName string, objectName string) ([]Object, error) {
	if orgName == "" {
		slog.Warn("orgName is empty")
		return nil, ErrOrgNameEmpty
	}

	if objectName == "" {
		slog.Warn("objectName is empty")
		return nil, ErrObjectNameEmpty
	}

	allowedNames, err := s.GetObjectsNameByOrgName(ctx, orgName)
	if err != nil {
		//slog.Error("GetObjectsNameByOrgName err in GetObjectsByOrgNameAndObjectName", "err", err)
		return nil, err
	}

	allowedSet := make(map[string]bool, len(allowedNames))
	for _, name := range allowedNames {
		allowedSet[name] = true
	}

	if !allowedSet[objectName] {
		slog.Warn("Object name not allowed", "name", objectName)
		return nil, fmt.Errorf("%w: %s", ErrObjectNameNotAllowed, objectName)
	}

	objects, err := s.repo.FindObjectsByOrgNameAndObjectName(ctx, orgName, objectName)
	if err != nil {
		slog.Error("FindObjectsByOrgNameAndObjectName err", "err", err)
		return nil, err
	}

	return objects, nil
}
