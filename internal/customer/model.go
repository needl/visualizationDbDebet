package customer

type Customer struct {
	//ID   string    `json:"id"`
	Name string `json:"name"`
}

type Summary struct {
	ContractorsCount    int     `db:"contractors_count" json:"contractors_count"`
	TotalDebet          float64 `db:"total_debet" json:"total_debet"`
	TotalDebetLong      float64 `db:"total_debet_long" json:"total_debet_long"`
	TotalDebetOverdue   float64 `db:"total_debet_overdue" json:"total_debet_overdue"`
	TotalContractAmount float64 `db:"total_contract_amount" json:"total_contract_amount"`
	TotalPaidAmount     float64 `db:"total_paid_amount" json:"total_paid_amount"`
	TotalAcceptedAmount float64 `db:"total_accepted_amount" json:"total_accepted_amount"`
	Percentage          float64 `json:"percentage"`
}

type TopItem struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type BlockFactors struct {
	BankruptcyCount                *int `db:"bankrot_count" json:"bankrot_count"`
	LiquidationCount               *int `db:"likvidatsiya_count" json:"likvidatsiya_count"`
	UnreliabilityCount             *int `db:"nedostovernost_count" json:"nedostovernost_count"`
	ExcludedCount                  *int `db:"isklyuchenie_count" json:"isklyuchenie_count"`
	ForeignAgentCount              *int `db:"inostrannye_count" json:"inostrannye_count"`
	ExtremeTerrCount               *int `db:"eks_ter_count" json:"eks_ter_count"`
	RegistryOfUnscrupulous         *int `db:"nedobrosovestn_count" json:"nedobrosovestn_count"`
	AdministrativeResponsibility   *int `db:"admin_otvet_count" json:"admin_otvet_count"`
	IntentionBankrupt              *int `db:"nam_bankrot_count" json:"nam_bankrot_count"`
	AccountBlockingCount           *int `db:"blokirovka_count" json:"blokirovka_count"`
	AvgWorkersListLessThanOneCount *int `db:"chisl_count" json:"chisl_count"`
}
