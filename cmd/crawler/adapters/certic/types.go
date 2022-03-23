package certic

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	Balances []struct {
		Denom  string          `json:"denom"`
		Amount decimal.Decimal `json:"amount"`
	} `json:"balances"`
}
