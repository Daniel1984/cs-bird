package aion

import "github.com/shopspring/decimal"

type TheoanReport struct {
	Content []struct {
		Address             string          `json:"address"`
		Balance             decimal.Decimal `json:"balance"`
		HasInternalTransfer bool            `json:"hasInternalTransfer"`
		Contract            bool            `json:"contract"`
		Tokens              []interface{}   `json:"tokens"`
		Nonce               int             `json:"nonce"`
		TransactionHash     string          `json:"transactionHash"`
		LastBlockNumber     int             `json:"lastBlockNumber"`
	} `json:"content"`
}
