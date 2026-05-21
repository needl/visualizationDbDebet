package contractoranalysis

import "time"

type Summary struct {
	ContractsSum        float64 `db:"contracts_sum" json:"contracts_sum"`
	ObjectsCount        int     `db:"objects_count" json:"objects_count"`
	OverdueObjectsCount int     `db:"overdue_objects_count" json:"overdue_objects_count"`
}

type ObjectNode struct {
	ObjectName        string  `db:"object_name" json:"object_name"`
	ContractSum       float64 `db:"contract_sum" json:"contract_sum"`
	ReadinessPercent  *string `db:"readiness_percent" json:"readiness_percent"`
	RiskLevel         string  `db:"risk_level" json:"risk_level"`
	CustomerName      string  `db:"customer_name" json:"customer_name"`
	OverdueDebtAmount float64 `db:"overdue_debt_amount" json:"overdue_debt_amount"`
}

type CustomerNode struct {
	CustomerName string       `db:"customer_name" json:"customer_name"`
	ObjectsCount int          `db:"objects_count" json:"objects_count"`
	Objects      []ObjectNode `db:"objects" json:"objects"`
}

type ObjectDetails struct {
	Status            string     `db:"status" json:"status"`
	CustomerName      string     `db:"customer_name" json:"customer_name"`
	ContractorName    string     `db:"contractor_name" json:"contractor_name"`
	ObjectName        string     `db:"object_name" json:"object_name"`
	ContractSum       float64    `db:"contract_sum" json:"contract_sum"`
	PaidSum           float64    `db:"paid_sum" json:"paid_sum"`
	ReadinessPercent  *string    `db:"readiness_percent" json:"readiness_percent"`
	TDCSum            float64    `db:"tdc_sum" json:"tdc_sum"`
	RVExists          bool       `db:"rv_exists" json:"rv_exists"`
	DebetSum          float64    `db:"debet_sum" json:"debet_sum"`
	OverdueDebtAmount float64    `db:"overdue_debt_amount" json:"overdue_debt_amount"`
	AcceptedPercent   float64    `db:"accepted_percent" json:"accepted_percent"`
	WorkStartDate     *time.Time `db:"work_start_date" json:"work_start_date"`
	WorkEndDate       *time.Time `db:"work_end_date" json:"work_end_date"`
}

type Analytics struct {
	ContractorName string         `db:"contractor_name" json:"contractor_name"`
	Summary        Summary        `db:"summary" json:"summary"`
	Customers      []CustomerNode `db:"customers" json:"customers"`
	SelectedObject *ObjectDetails `db:"selected_object" json:"selected_object,omitempty"`
}
