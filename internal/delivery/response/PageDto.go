package response

type PageDto struct {
	ID                *int    `json:"id"`
	CountSourceOrg    int     `db:"count_source_org" json:"count_source_org"`       // Количество заказчиков из дебет // Уникальные
	CountContracts    int     `db:"count_contracts" json:"count_contracts"`         // Количество контрактов //Уникальные
	SumContractAmount float64 `db:"sum_contract_amount" json:"sum_contract_amount"` // Сумма контрактов
	SumDebetTotal     float64 `db:"sum_debet_total" json:"sum_debet_total"`         // Сумма ДЗ debet.debt_2025_12_31_total
	SumDebetOverdue   float64 `db:"sum_debet_overdue" json:"sum_debet_overdue"`     // Просроченная задолженность
}
