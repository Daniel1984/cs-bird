package dfinity

import "github.com/shopspring/decimal"

type NetworkIdentifier struct {
	Blockchain string `json:"blockchain"`
	Network    string `json:"network"`
}

type AccountIdentifier struct {
	Address string `json:"address"`
}

type BalanceReport struct {
	Balances []struct {
		Value decimal.Decimal `json:"value"`
	} `json:"balances"`
}
