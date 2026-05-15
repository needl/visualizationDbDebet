package response

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"visualizationDbDebet/internal/domainconst"
)

// Repository отвечает за подключение к бд и работу с ней
type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetPageDto(ctx context.Context) (*Response, error) {
	var pageDto Response

	query := `
		select
			count(distinct source_org_name) as count_source_org,
			count(distinct contract_number) as count_contracts,
			round(coalesce(sum(contract_amount), 0)) as sum_contract_amount,
			round(coalesce(sum(debt_2025_12_31_total), 0)) as sum_debet_total,
			round(coalesce(sum(debt_2025_12_31_overdue), 0)) as sum_debet_overdue
		from debet
		where source_org_name != $1
	`

	if err := r.db.GetContext(ctx, &pageDto, query, domainconst.ExcludedSourceOrgName); err != nil {
		return nil, fmt.Errorf("failed to get dto without in response repo: %w", err)
	}
	return &pageDto, nil
}

func (r *Repository) GetPageDtoWithMIP(ctx context.Context) (*Response, error) {
	var pageDto Response
	query := `
		select
			count(distinct source_org_name) as count_source_org,
			count(distinct contract_number) as count_contracts,
			round(coalesce(sum(contract_amount), 0)) as sum_contract_amount,
			round(coalesce(sum(debt_2025_12_31_total), 0)) as sum_debet_total,
			round(coalesce(sum(debt_2025_12_31_overdue), 0)) as sum_debet_overdue
		from debet
	`

	if err := r.db.GetContext(ctx, &pageDto, query); err != nil {
		return nil, fmt.Errorf("failed to get dto in response repo: %w", err)
	}
	return &pageDto, nil
}
