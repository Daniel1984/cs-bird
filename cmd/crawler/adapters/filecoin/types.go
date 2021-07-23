package filecoin

import "github.com/shopspring/decimal"

type Message struct {
	CID       string `json:"cid"`
	Height    int    `json:"height"`
	Timestamp int    `json:"timestamp"`
	From      string `json:"from"`
	To        string `json:"to"`
	Nonce     int    `json:"nonce"`
	Value     decimal.Decimal
}

type MessagesReport struct {
	TotalCount int       `json:"totalCount"`
	Messages   []Message `json:"messages"`
}

type Balance struct {
	Balance decimal.Decimal `json:"balance"`
}

type BalanceReport []Balance
