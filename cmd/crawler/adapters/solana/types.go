package solana

import "github.com/shopspring/decimal"

type SolanaReport struct {
	Result struct {
		Value struct {
			Lamports decimal.Decimal `json:"lamports"`
		} `json:"value"`
	} `json:"result"`
}
