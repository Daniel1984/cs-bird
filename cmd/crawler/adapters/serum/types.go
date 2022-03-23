package serum

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	Succcess bool `json:"succcess"`
	Data     struct {
		TokenInfo struct {
			TokenAmount struct {
				Amount   decimal.Decimal `json:"amount"`
				Decimals int             `json:"decimals"`
				UIAmount decimal.Decimal `json:"uiAmount"`
			} `json:"tokenAmount"`
		} `json:"tokenInfo"`
	} `json:"data"`
}
