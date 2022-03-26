package luna

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	Balance []struct {
		Denom     string          `json:"denom"`
		Available decimal.Decimal `json:"available"`
	} `json:"balance"`
}
