package waves

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	Address    string          `json:"address"`
	Regular    decimal.Decimal `json:"regular"`
	Generating decimal.Decimal `json:"generating"`
	Available  decimal.Decimal `json:"available"`
	Effective  decimal.Decimal `json:"effective"`
}
