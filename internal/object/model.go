package object

import "time"

type Object struct {
	Name              *string    `db:"construction_object" json:"construction_object"`
	Contractor        *string    `db:"counterparty_name" json:"counterparty_name"`
	WorkStartDate     *time.Time `db:"work_start_date" json:"work_start_date"`
	WorkEndDate       *time.Time `db:"work_end_date" json:"work_end_date"`
	BuildReadyPercent *string    `db:"build_ready_percent" json:"build_ready_percent"`
	PermissionToEnter *bool      `db:"permission_to_enter" json:"permission_to_enter"`
	ConclusionMke     *bool      `db:"conclusion" json:"conclusion"`
	HardContractPrice *float64   `db:"hard_contract_price" json:"hard_contract_price"`
	ContractAmount    *float64   `db:"contract_amount" json:"contract_amount"`
	PaidAmount        *float64   `db:"paid_amount" json:"paid_amount"`
	AcceptedAmount    *float64   `db:"accepted_amount" json:"accepted_amount"`
	DebetTotal2024    *float64   `db:"debt_2024_12_31_total" json:"debt_2024_12_31_total"`
	DebetOverdue2024  *float64   `db:"debt_2024_12_31_overdue" json:"debt_2024_12_31_overdue"`
	DebetTotal2025    *float64   `db:"debt_2026_03_31_total" json:"debt_2026_03_31_total"`
	DebetOverdue2025  *float64   `db:"debt_2026_03_31_overdue" json:"debt_2026_03_31_overdue"`
}
