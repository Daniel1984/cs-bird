package cosmos

import "github.com/shopspring/decimal"

type CosmosstationReport struct {
	Balances []struct {
		Amount decimal.Decimal `json:"amount"`
	}
}

type AtomscanReport struct {
	Height string `json:"height"`
	Result []struct {
		Denom  string          `json:"denom"`
		Amount decimal.Decimal `json:"amount"`
	} `json:"result"`
}
