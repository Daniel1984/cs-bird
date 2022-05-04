package elrond

import "github.com/shopspring/decimal"

type Explorer1Report struct {
	TxCount int             `json:"txCount"`
	Nonce   int64           `json:"nonce"`
	Balance decimal.Decimal `json:"balance"`
}

type Explorer2Report struct {
	TxCount int             `json:"txCount"`
	Nonce   int64           `json:"nonce"`
	Balance decimal.Decimal `json:"balance"`
}

type Explorer3Report struct {
	TxCount int             `json:"txCount"`
	Nonce   int64           `json:"nonce"`
	Balance decimal.Decimal `json:"balance"`
}
