package debet

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllDebet(ctx context.Context) ([]Debet, error) {
	var debets []Debet

	query := `
			select source_org_name,
			       counterparty_name,
					contract_number,
					contract_date,
					contract_amount,
					construction_object,
					debt_2025_12_31_total,
					debt_2025_12_31_overdue,
					construction_title
				from debet
			where source_org_name not like '%Мосинж%'
			`

	if err := r.db.SelectContext(ctx, &debets, query); err != nil {
		return nil, err
	}

	return debets, nil
}

func (r *Repository) GetByOrgName(ctx context.Context, orgName string) (*Debet, error) {
	var debet Debet

	query := `
				select source_org_name,
			       		counterparty_name,
						contract_number,
						contract_date,
						contract_amount,
						construction_object,
						debt_2025_12_31_total,
						debt_2025_12_31_overdue,
						construction_title
					from debet
				where source_org_name = $1
			`

	if err := r.db.GetContext(ctx, &debet, query, orgName); err != nil {
		return nil, err
	}

	return &debet, nil
}
