package types

import (
	"github.com/shopspring/decimal"
)

type Checkpoint struct {
	Height  int
	Coin    string
	Address string
	Balance decimal.Decimal
}
