package near

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	Result struct {
		Amount      decimal.Decimal `json:"amount"`
		BlockHash   string          `json:"block_hash"`
		BlockHeight int             `json:"block_height"`
	} `json:"result"`
	ID string `json:"id"`
}
