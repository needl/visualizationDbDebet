package debet

import (
	"context"
	"database/sql"
	"errors"
	"visualizationBdDebet/internal/domainconst"

	"github.com/jmoiron/sqlx"
)

// Repository отвечает за подключение к бд и работу с ней
type Repository struct {
	db *sqlx.DB
}

// NewRepository конструктор репозитория, который возвращает конкретный
// инстанс репы с просеченым коннектом
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// GetAllView отвечает за обращение к базе для получения списка
// всех сущностей debet без учёта Мосинжпроекта
// ctx используется для отмены, таймаутов и тп
func (r *Repository) GetAllView(ctx context.Context) ([]View, error) {
	var debets []View

	query := `
		select
			id,
			source_org_name,
			counterparty_name,
			contract_number,
			contract_date,
			contract_amount,
			construction_object,
			debt_2025_12_31_total,
			debt_2024_12_31_total,
			debt_2025_12_31_overdue,
			debt_2024_12_31_overdue,
			construction_title
		from debet
		where source_org_name != $1
		order by id
	`

	if err := r.db.SelectContext(ctx, &debets, query, domainconst.ExcludedSourceOrgName); err != nil {
		// Непредвиденная ошибка при обращении к базе
		return nil, err
	}

	return debets, nil

}

func (r *Repository) GetAllViewWithMIP(ctx context.Context) ([]View, error) {
	var debets []View

	query := `
		select
			id,
			source_org_name,
			counterparty_name,
			contract_number,
			contract_date,
			contract_amount,
			construction_object,
			debt_2025_12_31_total,
			debt_2025_12_31_overdue,
			construction_title
		from debet
		order by id
	`

	if err := r.db.SelectContext(ctx, &debets, query); err != nil {
		// Непредвиденная ошибка при обращении к базе
		return nil, err
	}

	return debets, nil

}

// GetViewByOrgName отвечает за обращение к базе для получения конкретной
// сущности debet по его названию(orgName)
// ctx используется для отмены, таймаутов и тп
func (r *Repository) GetViewByOrgName(ctx context.Context, orgName string) (*View, error) {
	var debet View

	query := `
		select
			source_org_name,
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &debet, nil
}
