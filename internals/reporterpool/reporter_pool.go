package reporterpool

import (
	"github.com/cs-bird/internals/randnum"
	"github.com/shopspring/decimal"
)

type Reporter func(string) (decimal.Decimal, error)

// PullRandReporter accepts reporter pool, pulls random reporter from it,
// and returns remainder of the pool and pulled reporter
func PullRandReporter(repPool []Reporter) ([]Reporter, Reporter) {
	randIdx := randnum.InRange(0, int64(len(repPool)))
	rep := repPool[randIdx]
	return append(repPool[:randIdx], repPool[randIdx+1:]...), rep
}
