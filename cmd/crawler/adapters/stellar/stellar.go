package stellar

import (
	"errors"
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "GAWPTHY6233GRWZZ7JXDMVXDUDCVQVVQ2SXCSTG3R3CNP5LQPDAHNBKL"
	explorerOneURL = "https://horizon.stellar.org/accounts/%s"
)

type Stellar struct {
	Name    string
	Address string
}

func New() *Stellar {
	return &Stellar{
		Name:    "stellar",
		Address: address,
	}
}

func (e *Stellar) Get() (cp models.Checkpoint, err error) {
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
		return decimal.NewFromInt(0), err
	}

	for _, v := range rep.Balances {
		if v.AssetType == "native" {
			return v.Balance, nil
		}
	}

	return decimal.NewFromInt(0), errors.New("stellar balance not found")
}
