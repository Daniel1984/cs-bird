package omni

import "github.com/shopspring/decimal"

type OmniReporterOne struct {
	Balance []struct {
		Symbol   string          `json:"symbol"`
		Value    decimal.Decimal `json:"value"`
		Error    bool            `json:"error,omitempty"`
		Errormsg string          `json:"errormsg,omitempty"`
	} `json:"balance"`
}
