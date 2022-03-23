package vsys

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Address    string          `json:"Address"`
		Regular    string          `json:"regular"`
		RegularRaw decimal.Decimal `json:"regularRaw"`
	} `json:"data"`
}
