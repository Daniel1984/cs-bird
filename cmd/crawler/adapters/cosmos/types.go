package cosmos

import "github.com/shopspring/decimal"

type BalanceReport struct {
	Balances []struct {
		Amount decimal.Decimal `json:"amount"`
	}
}
