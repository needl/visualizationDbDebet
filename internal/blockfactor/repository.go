package blockfactor

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

func (r *Repository) GetViewAll(ctx context.Context) ([]View, error) {
	var blocks []View

	query := `
				select
					b.id,
					b.kod_nalogoplatelshchika,
					b.priznanie_bankrotom,
					b.likvidatsiya,
					b.nedostovernost_egryul,
					b.isklyuchenie_egryul,
					b.ekstremizm_terrorizm,
					b.reestr_nedobrosovestnyh_postavshchikov,
					b.administrativnaya_otvetstvennost_19_28,
					b.blokirovka_schetov,
					b.srednespisochnaya_chislennost_le_1
					from debet d
				left join blockfactor b
				on b.kod_nalogoplatelshchika = d.counterparty_inn
				order by b.id
	`

	if err := r.db.SelectContext(ctx, &blocks, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return blocks, nil
}

func (r *Repository) GetViewById(ctx context.Context, id int) (*View, error) {
	var block View

	query := `
				select
					id,
					kod_nalogoplatelshchika,
					priznanie_bankrotom,
					likvidatsiya,
					nedostovernost_egryul,
					isklyuchenie_egryul,
					ekstremizm_terrorizm,
					reestr_nedobrosovestnyh_postavshchikov,
					administrativnaya_otvetstvennost_19_28,
					blokirovka_schetov,
					srednespisochnaya_chislennost_le_1
					from blockfactor d
				where id = $1
	`

	if err := r.db.GetContext(ctx, &block, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &block, nil
}
