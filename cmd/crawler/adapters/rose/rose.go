package rose

import (
	"errors"
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address         = "oasis1qptlwu8arx7l8x0427f0xxke3swd27ctygukv0h5"
	oasismonitorURL = "https://api.oasismonitor.com/data/accounts/%s"
	oasisscanURL    = "https://www.oasisscan.com/mainnet/chain/account/info/%s"
)

type Rose struct {
	Name    string
	Address string
}

func New() *Rose {
	return &Rose{
		Name:    "rose",
		Address: address,
	}
}

func (c *Rose) Get() (cp types.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		oasismonitorReport,
		oasisscanReport,
	}

	for {
		var reporter reporterpool.Reporter
		reporterPool, reporter = reporterpool.PullRandReporter(reporterPool)

		balance, err := reporter(c.Address)
		if err == nil {
			cp.Balance = balance
			cp.Address = c.Address
			cp.Coin = c.Name
			break
		} else {
			// don't care about failure here since if 1 reporter fails, another might succeed
			fmt.Printf("failed getting %s wallet info: %s\n", c.Name, err)
		}

		if len(reporterPool) == 0 {
			err = errors.New("run out of connections for: " + c.Name)
			break
		}
	}

	return
}

func oasismonitorReport(address string) (decimal.Decimal, error) {
	rep := OasisMonitorReport{}
	req := request.
		New("GET", fmt.Sprintf(oasismonitorURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	formattedBalance := rep.TotalBalance.Div(decimal.NewFromInt(1000000000)).Round(4)
	return formattedBalance, nil
}

func oasisscanReport(address string) (decimal.Decimal, error) {
	rep := OasisScanReport{}
	req := request.
		New("GET", fmt.Sprintf(oasisscanURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Data.Total, nil
}
