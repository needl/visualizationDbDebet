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

func (s *Service) GetContractorsWithCurrDeb(ctx context.Context) ([]Debet, error) {
	contractors, err := s.repo.FindContractorsWithCurrDebet(ctx)
	if err != nil {
		slog.Error("FindContractorsWithCurrDebet err:", "err", err)
		return nil, err
	}

	if contractors == nil {
		return []Debet{}, nil
	}

	return contractors, nil
}

func (s *Service) GetContractorsWithOverdueDeb(ctx context.Context) ([]Debet, error) {
	contractors, err := s.repo.FindContractorsWithOverdueDebet(ctx)
	if err != nil {
		slog.Error("FindContractorsWithOverdueDebet err:", "err", err)
		return nil, err
	}

	if contractors == nil {
		return []Debet{}, nil
	}

	return contractors, nil
}

func (s *Service) GetContractorsWithBlockFactors(
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

	contractors, err := s.repo.FindContractorsWithBlockFactors(ctx, sourceOrgName, columnName)
	if err != nil {
		slog.Error("FindContractorsWithBlockFactors err:", "err", err)
		if errors.Is(err, ErrColumnNotAllowed) {
			return nil, apperr.NewInvalidArgument(err.Error())
		}
		return nil, err
	}

	return contractors, nil
}

func (s *Service) GetContractorsWithDebt(
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

	contractors, err := s.repo.FindContractorWithDebt(ctx, sourceOrgName, counterpartyName)
	if err != nil {
		slog.Error("FindContractorWithDebt err:", "err", err)
		return nil, err
	}

	return contractors, nil
}

func (s *Service) GetContractorsWithOverdue(
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

	contractors, err := s.repo.FindContractorWithOverdue(ctx, sourceOrgName, counterpartyName)
	if err != nil {
		slog.Error("FindContractorWithOverdue err:", "err", err)
		return nil, err
	}

	return contractors, nil
}

func (s *Service) GetContractorForTable(
	ctx context.Context,
	counterpartyName string,
) ([]Table, error) {

	if counterpartyName == "" {
		slog.Warn("counterpartyName is empty")
		return nil, apperr.NewInvalidArgument("counterpartyName is required")
	}

	contractors, err := s.repo.FindContractorForTable(ctx, counterpartyName)
	if err != nil {
		slog.Error("FindContractorForTable err:", "err", err)
		return nil, err
	}

	return contractors, nil
}
