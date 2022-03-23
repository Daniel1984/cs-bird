package tron

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	WithPriceTokens []struct {
		Amount    decimal.Decimal `json:"amount"`
		Balance   decimal.Decimal `json:"balance"`
		TokenAbbr string          `json:"tokenAbbr"`
		TokenName string          `json:"tokenName"`
	} `json:"withPriceTokens"`
	Balance decimal.Decimal `json:"balance"`
}
