package contractoranalysis

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

type treeRow struct {
	CustomerName     string          `db:"customer_name"`
	ObjectName       string          `db:"object_name"`
	ContractSum      float64         `db:"contract_sum"`
	ReadinessPercent sql.NullFloat64 `db:"readiness_percent"`
	OverdueDebtSum   float64         `db:"overdue_debt_sum"`
}

type objectDetailsRow struct {
	CustomerName     string          `db:"customer_name"`
	ContractSum      float64         `db:"contract_sum"`
	PaidSum          float64         `db:"paid_sum"`
	AcceptedSum      float64         `db:"accepted_sum"`
	DebetSum         float64         `db:"debet_sum"`
	OverdueDebtSum   float64         `db:"overdue_debt_sum"`
	ReadinessPercent sql.NullFloat64 `db:"readiness_percent"`
	WorkStartDate    sql.NullTime    `db:"work_start_date"`
	WorkEndDate      sql.NullTime    `db:"work_end_date"`
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) findContractorNames(ctx context.Context) ([]string, error) {
	var contractors []string

	query := `
		select distinct counterparty_name
		from debet
		where coalesce(counterparty_name, '') <> ''
		order by counterparty_name
	`

	if err := r.db.SelectContext(ctx, &contractors, query); err != nil {
		return nil, err
	}

	return contractors, nil
}

func (r *Repository) findSummaryByContractorName(ctx context.Context, contractorName string) (*Summary, error) {
	var summary Summary

	query := `
		select
			coalesce(sum(contract_amount), 0.0) as contracts_sum,
			count(distinct nullif(construction_object, '')) as objects_count,
			avg(build_ready_percent)::float8 as avg_readiness_percent,
			count(distinct case
				when coalesce(debt_2025_12_31_overdue, 0.0) > 0.0
					then nullif(construction_object, '')
			end) as overdue_objects_count
		from debet
		where counterparty_name = $1
	`

	if err := r.db.GetContext(ctx, &summary, query, contractorName); err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *Repository) findTreeByContractorName(ctx context.Context, contractorName string) ([]treeRow, error) {
	rows := make([]treeRow, 0)

	query := `
		select
			source_org_name as customer_name,
			construction_object as object_name,
			coalesce(sum(contract_amount), 0.0) as contract_sum,
			avg(build_ready_percent)::float8 as readiness_percent,
			coalesce(sum(debt_2025_12_31_overdue), 0.0) as overdue_debt_sum
		from debet
		where counterparty_name = $1
			and coalesce(construction_object, '') <> ''
		group by source_org_name, construction_object
		order by source_org_name asc, construction_object asc
	`

	if err := r.db.SelectContext(ctx, &rows, query, contractorName); err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *Repository) findObjectDetails(
	ctx context.Context,
	contractorName string,
	customerName string,
	objectName string,
) (*objectDetailsRow, error) {
	var row objectDetailsRow

	query := `
		select
			min(source_org_name) as customer_name,
			coalesce(sum(contract_amount), 0.0) as contract_sum,
			coalesce(sum(paid_amount), 0.0) as paid_sum,
			coalesce(sum(accepted_amount), 0.0) as accepted_sum,
			coalesce(sum(debt_2025_12_31_total), 0.0) as debet_sum,
			coalesce(sum(debt_2025_12_31_overdue), 0.0) as overdue_debt_sum,
			avg(build_ready_percent)::float8 as readiness_percent,
			min(work_start_date) as work_start_date,
			max(work_end_date) as work_end_date
		from debet
		where counterparty_name = $1
			and source_org_name = $2
			and construction_object = $3
	`

	if err := r.db.GetContext(ctx, &row, query, contractorName, customerName, objectName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &row, nil
}
