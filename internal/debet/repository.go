package debet

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

// GetAllDebet отвечает за обращение к базе для получения списка
// всех сущностей debet без учёта Мосинжпроекта
// ctx используется для отмены, таймаутов и тп
func (r *Repository) GetAllDebet(ctx context.Context) ([]Debet, error) {
	var debets []Debet

	query := `
			select  id, 
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
			where source_org_name not like '%Мосинж%'
			`

	// Проверка ошибки
	if err := r.db.SelectContext(ctx, &debets, query); err != nil {

		// Если возвращается ошибка sql.ErrNoRows -> данные не найдены в базе
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		// Непредвиденная ошибка при обращении к базе
		return nil, err
	}

	return debets, nil
}

// GetByOrgName отвечает за обращение к базе для получения конкретной
// сущности debet по его названию(orgName)
// ctx используется для отмены, таймаутов и тп
func (r *Repository) GetByOrgName(ctx context.Context, orgName string) (*Debet, error) {
	var debet Debet

	query := `
				select  source_org_name,
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
