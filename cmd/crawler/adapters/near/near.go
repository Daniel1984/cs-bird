package near

import (
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "383e50ea1a754ed3acd0d59116f221add87adb82559f31ca6d377f058fe83375"
	explorerOneURL = "https://explorer.near.org/_next/data/qmZoGxGkMm0Xv511MFKII/accounts/%s.json?id=%s"
)

type Near struct {
	Name    string
	Address string
}

func New() *Near {
	return &Near{
		Name:    "near",
		Address: address,
	}
}

func (e *Near) Get() (cp models.Checkpoint, err error) {
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
		New("GET", fmt.Sprintf(explorerOneURL, address, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.PageProps.Account.TotalBalance, nil
}
