package customer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) findAllCustomers(ctx context.Context) ([]Customer, error) {
	var customers []Customer

	query := `
		select distinct source_org_name as name
		from debet_new
		order by source_org_name
`

	if err := r.db.SelectContext(ctx, &customers, query); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *Repository) findSummaryByCustomerID(ctx context.Context, id string) (*Summary, error) {
	var summary Summary

	/*query := `
					select
	    				count(distinct counterparty_inn) as contractors_count,
	    				coalesce(sum(debt_2026_03_31_total), 0) as total_debet,
	    				coalesce(sum(debt_2026_03_31_overdue), 0) as total_debet_overdue,
	    				coalesce(sum(debt_2026_03_31_long_term), 0) as total_debet_long,
						coalesce(sum(contract_amount), 0) as total_contract_amount,
						coalesce(sum(paid_amount), 0) as total_paid_amount,
						coalesce(sum(accepted_amount), 0) as total_accepted_amount
					from debet_new
					where source_org_name = $1
	`*/

	query := `
		select
			count(distinct counterparty_inn) as contractors_count,
			coalesce(sum(debt_2026_03_31_total), 0) as total_debet,
			coalesce(sum(debt_2026_03_31_overdue), 0) as total_debet_overdue,
			coalesce(sum(contract_amount), 0) as total_contract_amount,
			coalesce(sum(paid_amount), 0) as total_paid_amount,
			coalesce(sum(accepted_amount), 0) as total_accepted_amount
		from debet_new
		where source_org_name = $1
`

	if err := r.db.GetContext(ctx, &summary, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &Summary{}, nil
		}
		return nil, err
	}

	return &summary, nil
}

func (r *Repository) findTopItemsByCustomerID(ctx context.Context, id string) ([]TopItem, error) {
	items := make([]TopItem, 0)

	query := `
		select
			counterparty_name as name,
			coalesce(sum(debt_2026_03_31_total), 0) as value
		from debet_new
		where source_org_name = $1
		group by counterparty_name
		having coalesce(sum(debt_2026_03_31_total), 0) > 0
		order by value desc
		limit 10
`

	if err := r.db.SelectContext(ctx, &items, query, id); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository) findTopItemsOverdueByCustomerID(ctx context.Context, id string) ([]TopItem, error) {
	items := make([]TopItem, 0)

	query := `
		select
			counterparty_name as name,
			coalesce(sum(debt_2026_03_31_overdue), 0) as value
		from debet_new
		where source_org_name = $1
		group by counterparty_name
		having coalesce(sum(debt_2026_03_31_overdue), 0) > 0
		order by value desc
		limit 10
`

	if err := r.db.SelectContext(ctx, &items, query, id); err != nil {
		return []TopItem{}, err
	}

	return items, nil
}

func (r *Repository) findCountBlockFactorsByCustomerID(ctx context.Context, id string) (*BlockFactors, error) {
	var factor BlockFactors

	query := `
		select
			sum(b.priznanie_bankrotom) as bankrot_count,
			sum(b.likvidatsiya) as likvidatsiya_count,
			sum(b.nedostovernost_egryul) as nedostovernost_count,
			sum(b.isklyuchenie_egryul) as isklyuchenie_count,
			sum(b.inostrannye_agenty) as inostrannye_count,
			sum(b.ekstremizm_terrorizm) as eks_ter_count,
			sum(b.reestr_nedobrosovestnyh_postavshchikov) as nedobrosovestn_count,
			sum(b.administrativnaya_otvetstvennost_19_28) as admin_otvet_count,
			sum(b.namerenie_bankrotstvo) as nam_bankrot_count,
			sum(b.blokirovka_schetov) as blokirovka_count,
			sum(b.srednespisochnaya_chislennost_le_1) as chisl_count,
			sum(b.priznanie_bankrotom)+
			sum(b.likvidatsiya)+
			sum(b.nedostovernost_egryul)+
			sum(b.isklyuchenie_egryul)+
			sum(b.inostrannye_agenty)+
			sum(b.ekstremizm_terrorizm)+
			sum(b.reestr_nedobrosovestnyh_postavshchikov)+
			sum(b.administrativnaya_otvetstvennost_19_28)+
			sum(b.namerenie_bankrotstvo)+sum(b.blokirovka_schetov)+
			sum(b.srednespisochnaya_chislennost_le_1) as all_risks_count
		from (
			select distinct counterparty_inn
			from debet_new
			where source_org_name = $1
		) as d
		inner join blockfactor b on b.kod_nalogoplatelshchika = d.counterparty_inn
	`

	if err := r.db.GetContext(ctx, &factor, query, id); err != nil {
		return nil, err
	}

	return &factor, nil
}
