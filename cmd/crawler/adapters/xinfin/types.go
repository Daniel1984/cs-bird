package xinfin

import (
	"github.com/shopspring/decimal"
)

type ExplorerOneResponse struct {
	Balance       decimal.Decimal `json:"balance"`
	BalanceNumber decimal.Decimal `json:"balanceNumber"`
	AccountName   string          `json:"accountName"`
}
