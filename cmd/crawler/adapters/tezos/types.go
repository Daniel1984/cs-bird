package tezos

import "github.com/shopspring/decimal"

type ReporterOne struct {
	Balance decimal.Decimal `json:"balance"`
}

type ReporterTwo struct {
	Balance decimal.Decimal `json:"total_balance"`
}
