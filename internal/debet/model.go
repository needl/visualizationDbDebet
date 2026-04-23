package debet

import "time"

type Debet struct {
	ID                            int        `db:"id" json:"id"`                                                         //
	Inn                           string     `db:"source_org_inn" json:"inn"`                                            // Инн организации
	OrgName                       string     `db:"source_org_name" json:"Застройщик"`                                    // Застройщик
	ContractorName                *string    `db:"counterparty_name" json:"Подрядчик"`                                   // Подрядчик
	ContractorInn                 *string    `db:"counterparty_inn" json:"Инн подрядчика"`                               // Инн подрядчик
	ContractNumber                *string    `db:"contract_number" json:"Номер Контракта"`                               // Номер контракта
	ContractDate                  *time.Time `db:"contract_date" json:"Дата контракта"`                                  // Дата контракта
	ContractAmount                *float64   `db:"contract_amount" json:"Сумма контракта"`                               // Сумма контракта
	PaidAmount                    *float64   `db:"paid_amount" json:"Сумма выплат"`                                      // Сумма выплат
	AcceptedAmount                *float64   `db:"accepted_amount" json:"Принятые выплаты"`                              // Принятые выплаты
	WorkStartDate                 *time.Time `db:"work_start_date" json:"Дата начала строительства"`                     // Дата начала строительства
	WorkStartEnd                  *time.Time `db:"work_start_end" json:"Дата окончания строительства"`                   // Дата окончания строительства
	ConstructionObject            *string    `db:"construction_object" json:"Наименование объекта"`                      // Наименование объекта получения
	ConstructionTitle             *string    `db:"construction_title" json:"Титул объекта"`                              // Титул объекта
	DsCode                        *string    `db:"ds_code" json:"Дс код"`                                                // Дс код
	UinObject                     *string    `db:"uin_object" json:"Юин объекта"`                                        // Юин объекта
	GuaranteeAmount               *float64   `db:"guarantee_amount" json:"Гарантия выплаты"`                             // Гарантия выплаты
	GuaranteeExpiry               *time.Time `db:"guarantee_expiry" json:"Дата гарантии"`                                // Дата гарантии
	AdvancedGaranteeAmount        *float64   `db:"advance_guarantee_amount" json:"Средняя цена гарантии"`                // Средняя цена гарантии
	AdvancedGaranteeExpiry        *time.Time `db:"advance_guarantee_expiry" json:"Средняя дата гарантии"`                // Средняя дата гарантии
	TaxScoringScore               *int       `db:"tax_scoring_score" json:"Сумма налогов по скорингу"`                   // Сумма налогов по скорингу
	TaxScoringDate                *time.Time `db:"tax_scoring_date" json:"Дата скоринга"`                                // Дата скоринга
	DebetTotal24                  *float64   `db:"debt_2024_12_31_total" json:"Дебеторская задолженность 24"`            // Дебеторская задолженность 24 (письмо)
	DebetLongTerm24               *float64   `db:"debt_2024_12_31_long_term" json:"Дебеторская 24"`                      //Деб 24
	DebetOverdose24               *float64   `db:"debt_2024_12_31_overdue" json:"Дебеторская переплата 24"`              // Дебеторская переплата 24
	DebetTotal25                  *float64   `db:"debt_2025_12_31_total" json:"Дебеторская задолженность 25"`            // Дебеторская задолженность 25 (письмо)
	DebetLongTerm25               *float64   `db:"debt_2025_12_31_long_term" json:"Дебеторская 25"`                      //Деб 25
	DebetOverdose25               *float64   `db:"debt_2025_12_31_overdue" json:"Дебеторская переплата 25"`              // Дебеторская переплата 25
	DebetIncreaseTotal            *float64   `db:"debt_increase_total" json:"Дебеторское увеличение всего"`              // Дебеторское увеличение всего
	DebetIncreaseAdvances         *float64   `db:"debt_increase_advances" json:"Дебеторское увеличение аванса"`          // Дебеторское увеличение всего
	DebetIncreasePenalties        *float64   `db:"debt_increase_penalties" json:"Дебеторское увеличение пени"`           // Дебеторское увеличение пени
	DebetDecreaseTotal            *float64   `db:"debt_decrease_total" json:"Дебеторское уменьшение всего"`              // Дебеторское уменьшение всего
	DebetDecreaseAcceptedWorks    *float64   `db:"debt_decrease_accepted_works" json:"Уменьшение принятых работ"`        // Уменьшение принятых работ
	DebetDecreaseReturnedAdvances *float64   `db:"debt_decrease_return_advances" json:"Уменьшение возвращённых авансов"` // Уменьшение возвращённых авансов
	DebetDecreasePenaltyPayment   *float64   `db:"debt_decrease_penalty_payment" json:"Уменьшение выплат пени"`          // Уменьшение выплат пени
	DebetTotal26                  *float64   `db:"debt_2026_01_31_total" json:"Дебеторская задолженность 26"`            // Дебеторская задолженность 26 (письмо)
	DebetLongTerm26               *float64   `db:"debt_2025_01_31_long_term" json:"Дебеторская 26"`                      // Деб 26
	DebetOverdose26               *float64   `db:"debt_2025_01_31_overdue" json:"Дебеторская переплата 26"`              // Дебеторская переплата 26
	DebetInPretension             *float64   `db:"debt_in_pretension" json:"Дебет в претензии"`                          // Дебет в претензии
	DebetInLitigation             *float64   `db:"debet_in_litigation" json:"Дебет в судебном процессе"`                 // Дебет в судебном процессе
	RepaymentDate                 *time.Time `db:"repayment_date" json:"Дата повторной оплаты"`                          // Дата повторной оплаты
	Notes                         *string    `db:"notes" json:"Заметки"`                                                 // Заметки
	ReportDate                    *time.Time `db:"report_date" json:"Дата отчёта"`                                       // Дата отчёта
	UploadDate                    *time.Time `db:"upload_date" json:"Дата обновления"`                                   // Дата обновления
	CreatedAt                     *time.Time `db:"created_at" json:"Дата создания записи"`                               // Дата создания записи
}

type View struct {
	ID                 int        `db:"id" json:"id"`                                           //
	OrgName            string     `db:"source_org_name" json:"source_org_name"`                 // Застройщик
	ContractorName     *string    `db:"counterparty_name" json:"counterparty_name"`             // Подрядчик
	ContractNumber     *string    `db:"contract_number" json:"contract_number"`                 // Номер контракта
	ContractDate       *time.Time `db:"contract_date" json:"contract_date"`                     // Дата контракта
	ContractAmount     *float64   `db:"contract_amount" json:"contract_amount"`                 // Сумма контракта
	ConstructionObject *string    `db:"construction_object" json:"construction_object"`         // Наименование объекта получения
	DebetTotal         *float64   `db:"debt_2025_12_31_total" json:"debt_2025_12_31_total"`     // Дебеторская задолженность (письмо)
	DebetTotal2024     *float64   `db:"debt_2024_12_31_total" json:"debt_2024_12_31_total"`     // Дебеторская задолженность (письмо)
	DebetOverdose      *float64   `db:"debt_2025_12_31_overdue" json:"debt_2025_12_31_overdue"` // Дебеторская задолженность(просроченная)
	DebetOverdose2024  *float64   `db:"debt_2024_12_31_overdue" json:"debt_2024_12_31_overdue"` // Дебеторская задолженность(просроченная)
	ConstructionTitle  *string    `db:"construction_title" json:"construction_title"`           // Титул
}
