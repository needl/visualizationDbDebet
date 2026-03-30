package blockfactor

import (
	"database/sql/driver"
	"fmt"
)

type BlockFactor struct {
	ID                                 *int    `db:"id" json:"id"`
	Name                               *string `db:"naimenovanie" json:"name"`
	Inn                                *string `db:"kod_nalogoplatelshchika" json:"inn"`
	IsBankrupt                         boolInt `db:"priznanie_bankrotom" json:"is_bankrupt"`
	IsLiquidation                      boolInt `db:"likvidatsiya" json:"is_liquidation"`
	IsUnreliabilityEGRUL               boolInt `db:"nedostovernost_egryul" json:"is_unreliability_egrul"`
	IsExcludedEGRUL                    boolInt `db:"isklyuchenie_egryul" json:"is_excluded_egrul"`
	IsForeignAgent                     boolInt `db:"inostrannye_agenty" json:"is_foreign_agent"`
	IsExtremTerr                       boolInt `db:"ekstremizm_terrorizm" json:"is_extrem_terr"`
	IsRegistryOfUnscrupulous           boolInt `db:"reestr_nedobrosovestnyh_postavshchikov" json:"Is_registry_of_unscrupulous"`
	IsAdministrativeResponsibility1928 boolInt `db:"administrativnaya_otvetstvennost_19_28" json:"is_administrative_responsibility_1928"`
	IsIntentionBankrupt                boolInt `db:"namerenie_bankrotstvo" json:"is_intention_bankrupt"`
	IsAccountBlocking                  boolInt `db:"blokirovka_schetov" json:"is_account_blocking"`
	IsAvgWorkersListLessThanOne        boolInt `db:"srednespisochnaya_chislennost_le_1" json:"is_avg_workers_list_less_than_one"`
}

type View struct {
	ID                                 *int    `db:"id" json:"id"`
	Inn                                *string `db:"kod_nalogoplatelshchika" json:"inn"`
	IsBankrupt                         boolInt `db:"priznanie_bankrotom" json:"is_bankrupt"`
	IsLiquidation                      boolInt `db:"likvidatsiya" json:"is_liquidation"`
	IsUnreliabilityEGRUL               boolInt `db:"nedostovernost_egryul" json:"is_unreliability_egrul"`
	IsExcludedEGRUL                    boolInt `db:"isklyuchenie_egryul" json:"is_excluded_egrul"`
	IsForeignAgent                     boolInt `db:"inostrannye_agenty" json:"is_foreign_agent"`
	IsExtremeTerr                      boolInt `db:"ekstremizm_terrorizm" json:"is_extrem_terr"`
	IsRegistryOfUnscrupulous           boolInt `db:"reestr_nedobrosovestnyh_postavshchikov" json:"Is_registry_of_unscrupulous"`
	IsAdministrativeResponsibility1928 boolInt `db:"administrativnaya_otvetstvennost_19_28" json:"is_administrative_responsibility_1928"`
	IsIntentionBankrupt                boolInt `db:"namerenie_bankrotstvo" json:"is_intention_bankrupt"`
	IsAccountBlocking                  boolInt `db:"blokirovka_schetov" json:"is_account_blocking"`
	IsAvgWorkersListLessThanOne        boolInt `db:"srednespisochnaya_chislennost_le_1" json:"is_avg_workers_list_less_than_one"`
}

type boolInt bool

func (b *boolInt) Scan(value any) error {
	if value == nil {
		*b = false
		return nil
	}
	switch v := value.(type) {
	case int64:
		*b = v != 0
	case int:
		*b = v != 0
	case bool:
		*b = boolInt(v)
	default:
		return fmt.Errorf("cannot convert %T to boolInt", value)
	}
	return nil
}

func (b *boolInt) Value() (driver.Value, error) {
	if *b {
		return int64(1), nil
	} else {
		return int64(0), nil
	}
}

func (b *boolInt) MarshalJSON() ([]byte, error) {
	if *b {
		return []byte("true"), nil
	} else {
		return []byte("false"), nil
	}
}

/*func (b *boolInt) UnmarshalJSON(data []byte) error {
	var v bool
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*b = BoolInt(v)
	return nil
}*/
