package contractor

import (
	"context"
	"errors"
	"log/slog"
	"visualizationDbDebet/internal/apperr"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getContractorsWithCurrDeb(ctx context.Context) ([]Debet, error) {
	contractors, err := s.repo.findContractorsWithCurrDebet(ctx)
	if err != nil {
		slog.Error("FindContractorsWithCurrDebet err:", "err", err)
		return nil, err
	}

	if contractors == nil {
		return []Debet{}, nil
	}

	return contractors, nil
}

func (s *Service) getContractorsWithOverdueDeb(ctx context.Context) ([]Debet, error) {
	contractors, err := s.repo.findContractorsWithOverdueDebet(ctx)
	if err != nil {
		slog.Error("FindContractorsWithOverdueDebet err:", "err", err)
		return nil, err
	}

	if contractors == nil {
		return []Debet{}, nil
	}

	return contractors, nil
}

func (s *Service) getContractorsWithBlockFactors(
	ctx context.Context,
	sourceOrgName string,
	columnName string,
) ([]View, error) {

	if sourceOrgName == "" {
		slog.Warn("sourceOrgName is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	if columnName == "" {
		slog.Warn("columnName is empty")
		return nil, apperr.NewInvalidArgument("columnName is required")
	}

	contractors, err := s.repo.findContractorsWithBlockFactors(ctx, sourceOrgName, columnName)
	if err != nil {
		slog.Error("FindContractorsWithBlockFactors err:", "err", err)
		if errors.Is(err, errColumnNotAllowed) {
			return nil, apperr.NewInvalidArgument(err.Error())
		}
		return nil, err
	}

	return contractors, nil
}

func (s *Service) getContractorsWithDebt(
	ctx context.Context,
	sourceOrgName string,
	counterpartyName string,
) ([]Contractor, error) {

	if sourceOrgName == "" {
		slog.Warn("sourceOrgName is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	if counterpartyName == "" {
		slog.Warn("counterpartyName is empty")
		return nil, apperr.NewInvalidArgument("counterpartyName is required")
	}

	contractors, err := s.repo.findContractorWithDebt(ctx, sourceOrgName, counterpartyName)
	if err != nil {
		slog.Error("FindContractorWithDebt err:", "err", err)
		return nil, err
	}

	return contractors, nil
}

func (s *Service) getContractorsWithOverdue(
	ctx context.Context,
	sourceOrgName string,
	counterpartyName string,
) ([]Contractor, error) {

	if sourceOrgName == "" {
		slog.Warn("sourceOrgName is empty")
		return nil, apperr.NewInvalidArgument("orgName is required")
	}

	if counterpartyName == "" {
		slog.Warn("counterpartyName is empty")
		return nil, apperr.NewInvalidArgument("counterpartyName is required")
	}

	contractors, err := s.repo.findContractorWithOverdue(ctx, sourceOrgName, counterpartyName)
	if err != nil {
		slog.Error("FindContractorWithOverdue err:", "err", err)
		return nil, err
	}

	return contractors, nil
}

func (s *Service) getContractorForTable(
	ctx context.Context,
	counterpartyName string,
) ([]Table, error) {

	if counterpartyName == "" {
		slog.Warn("counterpartyName is empty")
		return nil, apperr.NewInvalidArgument("counterpartyName is required")
	}

	contractors, err := s.repo.findContractorForTable(ctx, counterpartyName)
	if err != nil {
		slog.Error("FindContractorForTable err:", "err", err)
		return nil, err
	}

	return contractors, nil
}
