package elrond

import "github.com/shopspring/decimal"

type ElrondReport struct {
	TxCount int             `json:"txCount"`
	Nonce   int64           `json:"nonce"`
	Balance decimal.Decimal `json:"balance"`
}

type ElrondScanReport struct {
	TxCount int             `json:"txCount"`
	Nonce   int64           `json:"nonce"`
	Balance decimal.Decimal `json:"balance"`
}
