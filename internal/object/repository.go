package object

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

func (r *Repository) FindObjectsNameByOrgName(ctx context.Context, orgName string) ([]string, error) {
	var objectsName []string

	/*query := `select distinct construction_object from debet where source_org_name = $1`*/
	query := `select distinct coalesce(construction_object, '') as construction_object from debet where source_org_name = $1`

	if err := r.db.SelectContext(ctx, &objectsName, query, orgName); err != nil {
		return nil, err
	}

	return objectsName, nil
}

func (r *Repository) FindObjectsByOrgNameAndObjectName(ctx context.Context,
	orgName string,
	objectName string,
) ([]Object, error) {
	var objects []Object

	query := `
	select
		construction_object,
		contract_amount,
		counterparty_name,
		work_start_date,
		work_end_date,
		paid_amount,
		accepted_amount,
		debt_2024_12_31_total,
		debt_2024_12_31_overdue,
		debt_2025_12_31_total,
		debt_2025_12_31_overdue
	from debet
	where source_org_name = $1 and construction_object = $2
	`

	if err := r.db.SelectContext(ctx, &objects, query, orgName, objectName); err != nil {
		return nil, err
	}

	return objects, nil
}
