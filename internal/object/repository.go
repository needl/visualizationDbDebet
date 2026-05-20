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

func (r *Repository) findObjectsNameByOrgName(ctx context.Context, orgName string) ([]string, error) {
	var objectsName []string

	/*query := `select distinct construction_object from debet_new where source_org_name = $1`*/
	query := `
		select distinct coalesce(construction_object, '') as construction_object
		from debet_new
		where source_org_name = $1
	`

	if err := r.db.SelectContext(ctx, &objectsName, query, orgName); err != nil {
		return nil, err
	}

	return objectsName, nil
}

func (r *Repository) findObjectByName(ctx context.Context, name string) ([]Object, error) {
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
			debt_2026_03_31_total,
			debt_2026_03_31_overdue,
			construction_readiness_percent as build_ready_percent,
			coalesce(
				nullif(btrim(mge_status::text), '') is not null
				and lower(btrim(mge_status::text)) not in ('нет', 'false', '0', 'null', 'не получено', '-'),
				false
			) as conclusion,
			coalesce(
				nullif(btrim(rv_status::text), '') is not null
				and lower(btrim(rv_status::text)) not in ('нет', 'false', '0', 'null', 'не получено', '-'),
				false
			) as permission_to_enter,
			fixed_contract_price as hard_contract_price
		from debet_new
		where construction_object = $1
	`

	if err := r.db.SelectContext(ctx, &objects, query, name); err != nil {
		return nil, err
	}

	return objects, nil
}

func (r *Repository) findObjectsByOrgNameAndObjectName(ctx context.Context,
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
			debt_2026_03_31_total,
			debt_2026_03_31_overdue,
			construction_readiness_percent as build_ready_percent,
			coalesce(
				nullif(btrim(mge_status::text), '') is not null
				and lower(btrim(mge_status::text)) not in ('нет', 'false', '0', 'null', 'не получено', '-'),
				false
			) as conclusion,
			coalesce(
				nullif(btrim(rv_status::text), '') is not null
				and lower(btrim(rv_status::text)) not in ('нет', 'false', '0', 'null', 'не получено', '-'),
				false
			) as permission_to_enter,
			fixed_contract_price as hard_contract_price
		from debet_new
		where source_org_name = $1 and construction_object = $2
	`

	if err := r.db.SelectContext(ctx, &objects, query, orgName, objectName); err != nil {
		return nil, err
	}

	return objects, nil
}
