package wax

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	CoreLiquidBalance string `json:"core_liquid_balance"`
}

type ExplorerTwoResponse struct {
	Account string `json:"account"`
	Tokens  []struct {
		Symbol    string          `json:"symbol"`
		Precision int             `json:"precision"`
		Amount    decimal.Decimal `json:"amount"`
		Contract  string          `json:"contract"`
	} `json:"tokens"`
}
