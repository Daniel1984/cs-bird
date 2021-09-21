package avax

import "github.com/shopspring/decimal"

type AvaxReport struct {
	Addresses []struct {
		ChainID   string `json:"chainID"`
		Address   string `json:"address"`
		PublicKey string `json:"publicKey"`
		Assets    struct {
			Info struct {
				ID               string          `json:"id"`
				TransactionCount int64           `json:"transactionCount"`
				UtxoCount        int64           `json:"utxoCount"`
				Balance          decimal.Decimal `json:"balance"`
				TotalReceived    decimal.Decimal `json:"totalReceived"`
				TotalSent        decimal.Decimal `json:"totalSent"`
			} `json:"FvwEAhmxKfeiG8SnEvq42hc6whRyY3EFYAvebMqDNDGCgxN5Z"`
		} `json:"assets"`
	} `json:"addresses"`
}
