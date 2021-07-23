package elrond

import "github.com/shopspring/decimal"

type Report struct {
	TxCount int             `json:"txCount"`
	Nonce   int             `json:"nonce"`
	Balance decimal.Decimal `json:"balance"`
}
