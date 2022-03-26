package luna

import (
	"errors"
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "terra1t0an4m6t47rp3mj57rdfzw6dpd3lw8erxjppgw"
	explorerOneURL = "https://fcd.terra.dev/v1/bank/%s"
)

type Luna struct {
	Name    string
	Address string
}

func New() *Luna {
	return &Luna{
		Name:    "luna",
		Address: address,
	}
}

func (e *Luna) Get() (cp models.Checkpoint, err error) {
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

	for _, v := range rep.Balance {
		if v.Denom == "uluna" {
			return v.Available.Div(decimal.NewFromInt(1000000)).Round(6), nil
		}
	}

	return decimal.NewFromInt(0), errors.New("luna alance not found")
}
