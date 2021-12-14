package tfuel

import "github.com/shopspring/decimal"

// https://explorer.thetatoken.org/
type ReporterOne struct {
	Body struct {
		Balance struct {
			Thetawei decimal.Decimal `json:"thetawei"`
			Tfuelwei decimal.Decimal `json:"tfuelwei"`
		} `json:"balance"`
	} `json:"body"`
}
