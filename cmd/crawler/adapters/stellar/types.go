package stellar

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	Balances []struct {
		Balance   decimal.Decimal `json:"balance"`
		AssetType string          `json:"asset_type"`
		AssetCode string          `json:"asset_code,omitempty"`
	} `json:"balances"`
}
