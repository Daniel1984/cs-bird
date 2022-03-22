package radix

import "github.com/shopspring/decimal"

type RadixReporterOneReqPld struct {
	NetworkIdentifier struct {
		Network string `json:"network"`
	} `json:"network_identifier"`
	AccountIdentifier struct {
		Address string `json:"address"`
	} `json:"account_identifier"`
}

type RadixReport struct {
	AccountBalances struct {
		LiquidBalances []struct {
			Value           decimal.Decimal `json:"value"`
			TokenIdentifier struct {
				Rri string `json:"rri"`
			} `json:"token_identifier"`
		} `json:"liquid_balances"`
	} `json:"account_balances"`
}
