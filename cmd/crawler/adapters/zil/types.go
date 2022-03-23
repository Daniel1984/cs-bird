package zil

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	Hash    string          `json:"hash"`
	Balance decimal.Decimal `json:"balance"`
	TxCount int             `json:"txCount"`
}
