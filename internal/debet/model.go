package debet

import "time"

type Debet struct {
	ID                 int       `db:"id" json:"id"`                                           //
	OrgName            string    `db:"source_org_name" json:"Застройщик"`                      // Застройщик
	ContractorName     string    `db:"counterparty_name" json:"Подрядчик"`                     // Подрядчик
	ContractNumber     string    `db:"contract_number" json:"Номер Контракта"`                 // Номер контракта
	ContractDate       time.Time `db:"contract_date" json:"Дата контракта"`                    // Дата контракта
	ContractAmount     float64   `db:"contract_amount" json:"Сумма контракта"`                 // Сумма контракта
	ConstructionObject string    `db:"construction_object" json:"Наименование объекта"`        // Наименование объекта получения
	DebetTotal         float64   `db:"debt_2025_12_31_total" json:"Дебеторская задолженность"` // Дебеторская задолженность (письмо)
	DebetOverdose      float64   `db:"debt_2025_12_31_overdue" json:"Дебеторская переплата"`   // Дебеторская переплата?
	ConstructionTitle  string    `db:"construction_title" json:"Титул"`                        // Титул
}
