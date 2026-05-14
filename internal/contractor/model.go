// Package contractor отвечает за вывод информации в графике "блок факторы"
// Выводит информацию по контрактам с подрядчиками контрагента, которые имеют блок-факторы
package contractor

import "time"

type Contractor struct {
	Object       *string    `db:"object" json:"object"`               // Наименование объекта
	ContractDate *time.Time `db:"contract_date" json:"contract_date"` // Дата заключения контракта
	WorkEndDate  *time.Time `db:"work_end_date" json:"work_end_date"` // Дата окончания работ
	Number       *string    `db:"number" json:"number"`               // Номер контракта
	Amount       *float64   `db:"amount" json:"amount"`               // Сумма контракта
	DebetTotal   *float64   `db:"debet_total" json:"debet_total"`     // Сумма задолженности
	DebetOverdue *float64   `db:"debet_overdue" json:"debet_overdue"` // Сумма просроченной задолженности
}

type Debet struct {
	Name        *string  `db:"name" json:"name"`                 // Наименование подрядчика
	ContractSum *float64 `db:"contract_sum" json:"contract_sum"` // Сумма контракта
	PaidSum     *float64 `db:"paid_sum" json:"paid_sum"`         // Перечислено
	AcceptedSum *float64 `db:"accepted_sum" json:"accepted_sum"` // Сумма принятых работ
	DebetSum    *float64 `db:"debet_sum" json:"debet_sum"`       // Сумма текущей задолженности
}

type View struct {
	Name         *string  `db:"name" json:"name"`                   // Наименование подрядчика
	Amount       *float64 `db:"amount" json:"amount"`               // Сумма контракта
	DebetTotal   *float64 `db:"debet_total" json:"debet_total"`     // Сумма задолженности
	DebetOverdue *float64 `db:"debet_overdue" json:"debet_overdue"` // Сумма просроченной задолженности
}

type Table struct {
	OrgName       *string    `db:"org_name" json:"org_name"`
	Object        *string    `db:"object" json:"object"`                   // Наименование объекта
	WorkStartDate *time.Time `db:"work_start_date" json:"work_start_date"` // Дата начала работ
	WorkEndDate   *time.Time `db:"work_end_date" json:"work_end_date"`     // Дата окончания работ
	Number        *string    `db:"number" json:"number"`                   // Номер контракта
	Amount        *float64   `db:"amount" json:"amount"`                   // Сумма контракта
	DebetTotal    *float64   `db:"debet_total" json:"debet_total"`         // Сумма задолженности
	DebetOverdue  *float64   `db:"debet_overdue" json:"debet_overdue"`     // Сумма просроченной задолженности
}
