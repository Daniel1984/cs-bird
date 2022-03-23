package serum

import (
	"errors"
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "96kuAN8CTwSVmpBNyRaq6NuyxJ7aWwx7PSY7cwoANwHK"
	explorerOneURL = "https://api.solscan.io/account?address=%s"
)

type Serum struct {
	Name    string
	Address string
}

func New() *Serum {
	return &Serum{
		Name:    "serum",
		Address: address,
	}
}

func (e *Serum) Get() (cp models.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{reporterOne}

	for {
		var reporter reporterpool.Reporter
		reporterPool, reporter = reporterpool.PullRandReporter(reporterPool)

		balance, err := reporter(e.Address)
		if err == nil {
			cp.Balance = balance
			cp.Address = e.Address
			cp.Coin = e.Name
			break
		} else {
			// don't care about failure here since if 1 reporter fails, another might succeed
			fmt.Printf("failed getting %s wallet info: %s\n", e.Name, err)
		}

		if len(reporterPool) == 0 {
			break
		}
	}

	return
}

func reporterOne(address string) (decimal.Decimal, error) {
	rep := ExplorerOneResponse{}
	req := request.
		New("GET", fmt.Sprintf(explorerOneURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), fmt.Errorf("serum reporter 1 %w", err)
	}

	if !rep.Succcess {
		return decimal.NewFromInt(0), errors.New("serum reporter 1 request no success")
	}

	return rep.Data.TokenInfo.TokenAmount.UIAmount, nil
}
