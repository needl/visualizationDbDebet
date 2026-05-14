package contract

import "time"

type Contract struct {
	ID                         *int       `db:"id" json:"id"`
	RowNumber                  *int       `db:"row_number" json:"row_number"`                                     // Номер строки
	Title                      *string    `db:"titul" json:"titul"`                                               // Титул
	Payer                      *string    `db:"payer" json:"payer"`                                               // Платильщик
	ContractorInn              *string    `db:"contractor_inn" json:"contractor_inn"`                             // Инн подрядчика
	ContractorName             *string    `db:"contractor_name" json:"contractor_name"`                           // Наименование подрядчика
	ContractNumber             *string    `db:"contract_number" json:"contract_number"`                           // Номер конртакта
	ConclusionDate             *time.Time `db:"conclusion_date" json:"conclusion_date"`                           // Дата заключения
	ValidationDate             *time.Time `db:"validaty_date" json:"validation_date"`                             // Дата проверки
	ContractType               *string    `db:"contract_type" json:"contract_type"`                               // Тип контракта
	Status                     *string    `db:"status" json:"status"`                                             // Статус
	StatusChangeDate           *time.Time `db:"status_change_date" json:"status_change_date"`                     // Дата изменения статуса
	CurrentStatus              *string    `db:"current_status" json:"current_status"`                             // Текущий статус
	TdcAmount                  *float64   `db:"tdc_amount" json:"tdc_amount"`                                     // Сумма тдс
	ContractCost               *float64   `db:"contract_cost" json:"contract_cost"`                               // Сумма контракта
	PaymentStart               *float64   `db:"payment_start" json:"payment_start"`                               // Начальная сумма контракта
	Payment2024                *float64   `db:"payment_2024" json:"payment_2024"`                                 // Оплата 2024
	Payment2025                *float64   `db:"payment_2025" json:"payment_2025"`                                 // Оплата 2025
	Payment2026                *float64   `db:"payment_2026" json:"payment_2026"`                                 // Оплата 2026
	AdvanceStart               *float64   `db:"advance_start" json:"advance_start"`                               //
	Advance2024                *float64   `db:"advance_2024" json:"advance_2024"`                                 //
	Advance2025                *float64   `db:"advance_2025" json:"advance_2025"`                                 //
	Advance2026                *float64   `db:"advance_2026" json:"advance_2026"`                                 //
	AdvanceTotal               *float64   `db:"advance_total" json:"advance_total"`                               //
	AdvanceRepaymentAgreements *float64   `db:"advance_repayment_agreements" json:"advance_repayment_agreements"` //
	AdvanceRepaymentActs       *float64   `db:"advance_repayment_acts" json:"advance_repayment_acts"`             //
	AdvanceReturn              *float64   `db:"advance_return" json:"advance_return"`                             //
	UnpaidAdvance              *float64   `db:"unrepaid_advance" json:"unpaid_advance"`                           // Неоплаченный аванс (МГЗ)
	ExecutionFromStart         *float64   `db:"execution_from_start" json:"execution_from_start"`                 //
	LimitCostAip               *float64   `db:"limit_cost_aip" json:"limit_cost_aip"`                             //
	TitleYear                  *int       `db:"title_year" json:"title_year"`                                     //
	PaymentCurrentYear         *float64   `db:"current_year" json:"current_year"`                                 //
	PaymentYearPlusTwo         *float64   `db:"year_plus2" json:"payment_year_plus_two"`                          //
	PaymentYearPlusThree       *float64   `db:"year_plus3" json:"payment_year_plus_three"`                        //
	DsCode                     *string    `db:"ds_code" json:"ds_code"`                                           //
	ObjectName                 *string    `db:"object_name" json:"object_name"`                                   //
	ObjectState                *string    `db:"object_state" json:"object_state"`                                 //
	ObjectReadiness            *float64   `db:"object_readiness" json:"object_readiness"`                         // Готовность
	Industry                   *string    `db:"industry" json:"industry"`                                         // Отрасль
	Developer                  *string    `db:"developer" json:"developer"`                                       //
	RvFactDate                 *time.Time `db:"rv_fact_date" json:"rv_fact_date"`                                 //
	InitialPrice               *float64   `db:"initial_price" json:"initial_price"`                               //
	ExpectedCostMge            *float64   `db:"expected_cost_mge" json:"expected_cost_mge"`                       //
	ContractId                 *string    `db:"contract_id" json:"contract_id"`                                   //
	ReceivablesTotal           *float64   `db:"receivables_total" json:"receivables_total"`                       // Задолженность
	CreatedAt                  *time.Time `db:"created_at" json:"created_at"`                                     //
	LoadedAt                   *time.Time `db:"loaded_at" json:"loaded_at"`                                       //
}

type View struct {
	ID              *int     `db:"id" json:"id"`
	UnpaidAdvance   *float64 `db:"unrepaid_advance" json:"unpaid_advance"`   // Неоплаченный аванс (МГЗ)
	ContractCost    *float64 `db:"contract_cost" json:"contract_cost"`       // Сумма контракта
	ObjectName      *string  `db:"object_name" json:"object_name"`           //
	Status          *string  `db:"status" json:"status"`                     // Статус
	TdcAmount       *float64 `db:"tdc_amount" json:"tdc_amount"`             // Сумма тдц
	ObjectState     *string  `db:"object_state" json:"object_state"`         //
	ObjectReadiness *float64 `db:"object_readiness" json:"object_readiness"` // Готовность
}
