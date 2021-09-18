package rose

import (
	"time"

	"github.com/shopspring/decimal"
)

type OasisMonitorReport struct {
	Address                     string          `json:"address"`
	LiquidBalance               decimal.Decimal `json:"liquid_balance"`
	EscrowBalance               decimal.Decimal `json:"escrow_balance"`
	EscrowDebondingBalance      decimal.Decimal `json:"escrow_debonding_balance"`
	DelegationsBalance          decimal.Decimal `json:"delegations_balance"`
	DebondingDelegationsBalance decimal.Decimal `json:"debonding_delegations_balance"`
	SelfDelegationBalance       decimal.Decimal `json:"self_delegation_balance"`
	TotalBalance                decimal.Decimal `json:"total_balance"` // 712900065069270
	CreatedAt                   time.Time       `json:"created_at"`
	LastActive                  time.Time       `json:"last_active"`
	Nonce                       int64           `json:"nonce"`
	Type                        string          `json:"type"`
}

type OasisScanReport struct {
	Code int `json:"code"`
	Data struct {
		Rank      int64           `json:"rank"`
		Address   string          `json:"address"`
		Available decimal.Decimal `json:"available"`
		Escrow    decimal.Decimal `json:"escrow"`
		Debonding decimal.Decimal `json:"debonding"`
		Total     decimal.Decimal `json:"total"` // 712900.0651
		Nonce     int64           `json:"nonce"`
	} `json:"data"`
}
