package polkadot

import "github.com/shopspring/decimal"

type PolkadotReport struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Account struct {
			Address string          `json:"address"`
			Balance decimal.Decimal `json:"balance"`
		} `json:"account"`
	} `json:"data"`
}
