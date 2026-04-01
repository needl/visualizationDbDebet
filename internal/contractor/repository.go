package contractor

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// FindContractorsWithBlockFactors возвращает список подрядчиков с блок-факторами
// по заданному наименованию организации-источника и признаку банкротства
func (r *Repository) FindContractorsWithBlockFactors(
	ctx context.Context,
	sourceOrgName string,
	columnName string,
) ([]Contractor, error) {
	var contractors []Contractor

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
		return nil, errors.New("column " + columnName + " is not allowed")
	}

	query := fmt.Sprintf(`
		select distinct on (c.contract_number)
		d.counterparty_name as name,
		d.construction_object as object,
		d.contract_number as number,
		d.contract_amount as amount,
		d.debt_2025_12_31_total as debet_total,
		d.debt_2025_12_31_overdue as debet_overdue,
		d.contract_date as contract_date,
		d.work_end_date,
		c.status
	from debet d
	inner join contracts c on c.contract_number = d.contract_number
	inner join blockfactor b on d.counterparty_inn = b.kod_nalogoplatelshchika
	where d.source_org_name = $1
	  and d.contract_amount is not null
	  and b.%s = 1
	  and not (c.status = 'Расторгнут' and d.debt_2025_12_31_total = 0)
	  order by c.contract_number
`, columnName)

	if err := r.db.SelectContext(ctx, &contractors, query, sourceOrgName); err != nil {
		return nil, err
	}

	return contractors, nil
}
