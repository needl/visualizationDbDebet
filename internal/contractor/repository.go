package contractor

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"visualizationDbDebet/internal/domainconst"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

var errColumnNotAllowed = errors.New("column is not allowed")

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) findContractorWithDebt(
	ctx context.Context,
	sourceOrgName string,
	counterpartyName string,
) ([]Contractor, error) {
	var contractors []Contractor

	query := `
		select
			d.construction_object as object,
			d.contract_number as number,
			d.contract_amount as amount,
			d.debt_2025_12_31_total as debet_total,
			d.debt_2025_12_31_overdue as debet_overdue,
			d.contract_date as contract_date,
			d.work_end_date
		from debet d
		where d.source_org_name = $1
			and d.counterparty_name = $2
			and d.debt_2025_12_31_total <> 0
	`

	if err := r.db.SelectContext(ctx, &contractors, query, sourceOrgName, counterpartyName); err != nil {
		return nil, err
	}

	return contractors, nil
}

func (r *Repository) findContractorWithOverdue(
	ctx context.Context,
	sourceOrgName string,
	counterpartyName string,
) ([]Contractor, error) {
	var contractors []Contractor

	query := `
		select
			d.construction_object as object,
			d.contract_number as number,
			d.contract_amount as amount,
			d.debt_2025_12_31_total as debet_total,
			d.debt_2025_12_31_overdue as debet_overdue,
			d.contract_date as contract_date,
			d.work_end_date
		from debet d
		where d.source_org_name = $1
			and d.counterparty_name = $2
			and d.debt_2025_12_31_overdue <> 0
	`

	if err := r.db.SelectContext(ctx, &contractors, query, sourceOrgName, counterpartyName); err != nil {
		return nil, err
	}

	return contractors, nil
}

func (r *Repository) findContractorsWithCurrDebet(ctx context.Context) ([]Debet, error) {
	var contractors []Debet

	query := `
		select
			counterparty_name as name,
			coalesce(sum(contract_amount), 0.00) as contract_sum,
			coalesce(sum(paid_amount), 0.00) as paid_sum,
			coalesce(sum(accepted_amount), 0.00) as accepted_sum,
			coalesce(sum(debt_2025_12_31_total), 0.00)
				- coalesce(sum(debt_2025_12_31_overdue), 0.00) as debet_sum
		from debet
		where source_org_name != $1
		group by counterparty_name
		having coalesce(sum(debt_2025_12_31_total), 0)
			- coalesce(sum(debt_2025_12_31_overdue), 0.00) != 0.00
		order by debet_sum desc
	`

	if err := r.db.SelectContext(ctx, &contractors, query, domainconst.ExcludedSourceOrgName); err != nil {
		return nil, err
	}

	return contractors, nil
}

func (r *Repository) findContractorsWithOverdueDebet(ctx context.Context) ([]Debet, error) {
	var contractors []Debet

	query := `
		select
			counterparty_name as name,
			coalesce(sum(contract_amount), 0.00) as contract_sum,
			coalesce(sum(paid_amount), 0.00) as paid_sum,
			coalesce(sum(accepted_amount), 0.00) as accepted_sum,
			coalesce(sum(debt_2025_12_31_overdue), 0.00) as debet_sum
		from debet
		where source_org_name != $1
		group by counterparty_name
		having coalesce(sum(debt_2025_12_31_overdue), 0) != 0.00
		order by debet_sum desc
`

	if err := r.db.SelectContext(ctx, &contractors, query, domainconst.ExcludedSourceOrgName); err != nil {
		return nil, err
	}

	return contractors, nil
}

// FindContractorsWithBlockFactors возвращает список подрядчиков с блок-факторами
// по заданному наименованию организации-источника и признаку банкротства
func (r *Repository) findContractorsWithBlockFactors(
	ctx context.Context,
	sourceOrgName string,
	columnName string,
) ([]View, error) {
	var contractors []View

	allowedColumns := map[string]bool{
		"priznanie_bankrotom":                    true,
		"likvidatsiya":                           true,
		"nedostovernost_egryul":                  true,
		"isklyuchenie_egryul":                    true,
		"inostrannye_agenty":                     true,
		"ekstremizm_terrorizm":                   true,
		"reestr_nedobrosovestnyh_postavshchikov": true,
		"administrativnaya_otvetstvennost_19_28": true,
		"namerenie_bankrotstvo":                  true,
		"blokirovka_schetov":                     true,
		"srednespisochnaya_chislennost_le_1":     true,
	}

	if !allowedColumns[columnName] {
		slog.Warn("ColumnName not allowed")
		return nil, fmt.Errorf("%w: %s", errColumnNotAllowed, columnName)
	}

	query := fmt.Sprintf(`
		select
			min(d.counterparty_name) as name,
			sum(d.contract_amount) as amount,
			sum(d.debt_2025_12_31_total) as debet_total,
			sum(d.debt_2025_12_31_overdue) as debet_overdue
		from debet d
		inner join blockfactor b on d.counterparty_inn = b.kod_nalogoplatelshchika
		where d.source_org_name = $1
			and b.%s = 1
		group by d.counterparty_inn;
	`, columnName)

	if err := r.db.SelectContext(ctx, &contractors, query, sourceOrgName); err != nil {
		return nil, err
	}

	return contractors, nil
}

func (r *Repository) findContractorForTable(
	ctx context.Context,
	counterpartyName string,
) ([]Table, error) {
	var contractors []Table

	query := `
		select
			d.source_org_name as org_name,
			d.work_start_date,
			d.work_end_date,
			d.contract_number as number,
			d.contract_amount as amount,
			d.construction_object as object,
			d.debt_2025_12_31_total as debet_total,
			d.debt_2025_12_31_overdue as debet_overdue
		from debet d
		where d.counterparty_inn = (
			select d2.counterparty_inn
			from debet d2
			where d2.counterparty_name = $1
				and d2.source_org_name != $2
				and d2.counterparty_inn is not null
			limit 1
		)
			and d.source_org_name != $2
		order by d.source_org_name asc, d.work_start_date asc
	`

	if err := r.db.SelectContext(ctx, &contractors, query, counterpartyName, domainconst.ExcludedSourceOrgName); err != nil {
		return nil, err
	}

	return contractors, nil
}
