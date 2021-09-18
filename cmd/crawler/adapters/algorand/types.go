package algorand

import "github.com/shopspring/decimal"

type AlgoexplorerReport struct {
	Account struct {
		Address                     string          `json:"address"`
		Amount                      decimal.Decimal `json:"amount"`
		AmountWithoutPendingRewards decimal.Decimal `json:"amount-without-pending-rewards"`
		CreatedAtRound              int             `json:"created-at-round"`
		Deleted                     bool            `json:"deleted"`
		MinBalance                  int             `json:"min-balance"`
		PendingRewards              int             `json:"pending-rewards"`
		RewardBase                  int             `json:"reward-base"`
		Rewards                     decimal.Decimal `json:"rewards"`
		Round                       int             `json:"round"`
		SigType                     string          `json:"sig-type"`
		Status                      string          `json:"status"`
	} `json:"account"`
	CurrentRound int `json:"current-round"`
}

type PurestakeReport struct {
	Address                     string          `json:"address"`
	Amount                      decimal.Decimal `json:"amount"`
	AmountWithoutPendingRewards decimal.Decimal `json:"amount-without-pending-rewards"`
	AppsLocalState              []interface{}   `json:"apps-local-state"`
	AppsTotalSchema             struct {
		NumByteSlice int `json:"num-byte-slice"`
		NumUint      int `json:"num-uint"`
	} `json:"apps-total-schema"`
	Assets         []interface{}   `json:"assets"`
	CreatedApps    []interface{}   `json:"created-apps"`
	CreatedAssets  []interface{}   `json:"created-assets"`
	PendingRewards int             `json:"pending-rewards"`
	RewardBase     int             `json:"reward-base"`
	Rewards        decimal.Decimal `json:"rewards"`
	Round          int             `json:"round"`
	Status         string          `json:"status"`
}
