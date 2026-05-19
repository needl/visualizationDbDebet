package contract

import (
	"context"
	"database/sql"
	"errors"

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
// всех сущностей contract без учёта Мосинжпроекта
// ctx используется для отмены, таймаутов и тп
func (r *Repository) getAllView(ctx context.Context) ([]View, error) {
	var contracts []View

	query := `
		select
			c.id,
			c.unrepaid_advance,
			c.contract_cost,
			c.object_name,
			c.status,
			c.tdc_amount,
			c.object_state,
			c.object_readiness
		from debet d
		left join contracts c
			on c.titul = d.construction_title
			and c.contract_number = d.contract_number
		order by id
	`

	if err := r.db.SelectContext(ctx, &contracts, query); err != nil {
		return nil, err
	}

	return contracts, nil
}

// GetViewById отвечает за обращение к базе для получения конкретной
// сущности contract по его id(id)
// ctx используется для отмены, таймаутов и тп
func (r *Repository) getViewByID(ctx context.Context, id int) (*View, error) {
	var contract View

	query := `
		select
			id,
			unrepaid_advance,
			contract_cost,
			object_name,
			status,
			tdc_amount,
			object_state,
			object_readiness
		from contracts
		where id = $1
	`

	if err := r.db.GetContext(ctx, &contract, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &contract, nil
}
