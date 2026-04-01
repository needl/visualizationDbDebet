package contractor

import (
	"context"
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
		di := contractors[i].ContractDate
		dj := contractors[j].ContractDate

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
