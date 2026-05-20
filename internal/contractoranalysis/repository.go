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

type summaryRow struct {
	ContractsSum        float64         `db:"contracts_sum"`
	ObjectsCount        int             `db:"objects_count"`
	AvgReadinessPercent sql.NullFloat64 `db:"avg_readiness_percent"`
	OverdueObjectsCount int             `db:"overdue_objects_count"`
}

type treeRow struct {
	CustomerName     string          `db:"customer_name"`
	ObjectName       string          `db:"object_name"`
	ContractSum      float64         `db:"contract_sum"`
	ReadinessPercent sql.NullFloat64 `db:"readiness_percent"`
	OverdueDebtSum   float64         `db:"overdue_debt_sum"`
}

type objectDetailsRow struct {
	ContractSum      float64         `db:"contract_sum"`
	PaidSum          float64         `db:"paid_sum"`
	AcceptedSum      float64         `db:"accepted_sum"`
	TDCSum           float64         `db:"tdc_sum"`
	DebetSum         float64         `db:"debet_sum"`
	OverdueDebtSum   float64         `db:"overdue_debt_sum"`
	RVExists         bool            `db:"rv_exists"`
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
		from debet_new
		where coalesce(counterparty_name, '') <> ''
		order by counterparty_name
	`

	if err := r.db.SelectContext(ctx, &contractors, query); err != nil {
		return nil, err
	}

	return contractors, nil
}

func (r *Repository) findSummaryByContractorName(ctx context.Context, contractorName string) (*Summary, error) {
	var row summaryRow

	query := `
		select
			coalesce(sum(contract_amount), 0.0) as contracts_sum,
			count(distinct nullif(construction_object, '')) as objects_count,
			avg(construction_readiness_percent)::float8 as avg_readiness_percent,
			count(distinct case
				when coalesce(debt_2026_03_31_overdue, 0.0) > 0.0
					then nullif(construction_object, '')
			end) as overdue_objects_count
		from debet_new
		where counterparty_name = $1
	`

	if err := r.db.GetContext(ctx, &row, query, contractorName); err != nil {
		return nil, err
	}

	summary := &Summary{
		ContractsSum:        row.ContractsSum,
		ObjectsCount:        row.ObjectsCount,
		OverdueObjectsCount: row.OverdueObjectsCount,
	}

	if row.AvgReadinessPercent.Valid {
		summary.AvgReadinessPercent = new(row.AvgReadinessPercent.Float64)
	}

	return summary, nil
}

func (r *Repository) findContractorExists(ctx context.Context, contractorName string) (bool, error) {
	var exists bool

	query := `
		select exists(
			select 1
			from debet_new
			where counterparty_name = $1
		)
	`

	if err := r.db.GetContext(ctx, &exists, query, contractorName); err != nil {
		return false, err
	}

	return exists, nil
}

func (r *Repository) findTreeByContractorName(ctx context.Context, contractorName string) ([]treeRow, error) {
	rows := make([]treeRow, 0)

	query := `
		select
			coalesce(source_org_name, '') as customer_name,
			construction_object as object_name,
			coalesce(sum(contract_amount), 0.0) as contract_sum,
			avg(construction_readiness_percent)::float8 as readiness_percent,
			coalesce(sum(debt_2026_03_31_overdue), 0.0) as overdue_debt_sum
		from debet_new
		where counterparty_name = $1
			and coalesce(construction_object, '') <> ''
		group by coalesce(source_org_name, ''), construction_object
		order by coalesce(source_org_name, '') asc, construction_object asc
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
			coalesce(sum(contract_amount), 0.0) as contract_sum,
			coalesce(sum(paid_amount), 0.0) as paid_sum,
			coalesce(sum(accepted_amount), 0.0) as accepted_sum,
			coalesce(sum(fixed_contract_price), 0.0) as tdc_sum,
			coalesce(sum(debt_2026_03_31_total), 0.0) as debet_sum,
			coalesce(sum(debt_2026_03_31_overdue), 0.0) as overdue_debt_sum,
			coalesce(
				bool_or(
					nullif(btrim(rv_status::text), '') is not null
					and lower(btrim(rv_status::text)) not in ('нет', 'false', '0', 'null', 'не получено', '-')
				),
				false
			) as rv_exists,
			avg(construction_readiness_percent)::float8 as readiness_percent,
			min(work_start_date) as work_start_date,
			max(work_end_date) as work_end_date
		from debet_new
		where counterparty_name = $1
			and source_org_name = $2
			and construction_object = $3
		having count(*) > 0
	`

	if err := r.db.GetContext(ctx, &row, query, contractorName, customerName, objectName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &row, nil
}
