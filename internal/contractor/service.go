package contractor

import (
	"context"
	"errors"
	"log/slog"
	"sort"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetContractorsWithBlockFactors(
	ctx context.Context,
	sourceOrgName string,
	columnName string,
) ([]Contractor, error) {

	if sourceOrgName == "" {
		slog.Warn("sourceOrgName is empty")
		return nil, nil
	}

	if columnName == "" {
		slog.Warn("columnName is empty")
		return nil, nil
	}

	contractors, err := s.repo.FindContractorsWithBlockFactors(ctx, sourceOrgName, columnName)
	if err != nil {
		slog.Error("FindContractorsWithBlockFactors err:", "err", err)
		return nil, err
	}

	sort.Slice(contractors, func(i, j int) bool {
		di := contractors[i].WorkEndDate
		dj := contractors[j].WorkEndDate

		if di == nil && dj == nil {
			return false
		}
		if di == nil {
			return false
		}
		if dj == nil {
			return true
		}
		return di.Before(*dj)
	})

	return contractors, nil
}

func (s *Service) GetContractorsWithCurrDeb(ctx context.Context) ([]DebetContractor, error) {
	contractors, err := s.repo.FindContractorsWithCurrDebet(ctx)
	if err != nil {
		slog.Error("FindContractorsWithCurrDebet err:", "err", err)
	}

	if contractors == nil {
		return nil, errors.New("contractors is nil")
	}

	return contractors, nil
}

func (s *Service) GetContractorsWithOverdueDeb(ctx context.Context) ([]DebetContractor, error) {
	contractors, err := s.repo.FindContractorsWithOverdueDebet(ctx)
	if err != nil {
		slog.Error("FindContractorsWithOverdueDebet err:", "err", err)
	}

	if contractors == nil {
		return nil, errors.New("contractors is nil")
	}

	return contractors, nil
}
