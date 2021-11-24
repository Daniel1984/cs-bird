package aion

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address   = "0xa08fd235db5de8023721b5f9b6e29c0ee5253982ccdd423a1cd9f682ae745821"
	theoanURL = "https://mainnet-api.theoan.com/aion/dashboard/getAccountDetails?accountAddress=%s"
)

type Aion struct {
	Name    string
	Address string
}

func New() *Aion {
	return &Aion{
		Name:    "aion",
		Address: address,
	}
}

func (e *Aion) Get() (cp types.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{theoanReport}

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

func theoanReport(address string) (decimal.Decimal, error) {
	rep := TheoanReport{}
	req := request.
		New("GET", fmt.Sprintf(theoanURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	if len(rep.Content) == 0 {
		return decimal.NewFromInt(0), errors.New("aion no balances detected")
	}

	for _, v := range rep.Content {
		if strings.HasSuffix(address, v.Address) {
			return v.Balance, nil
		}
	}

	return decimal.NewFromInt(0), errors.New("aion hw address not found")
}
