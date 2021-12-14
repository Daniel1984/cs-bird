package tfuel

import (
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "0xDd3C4d1E564384414AF671d6d334521C79266BFD"
	explorerOneURL = "https://explorer.thetatoken.org:8443/api/account/update/%s"
)

type Tfuel struct {
	Name    string
	Address string
}

func New() *Tfuel {
	return &Tfuel{
		Name:    "tfuel",
		Address: address,
	}
}

func (e *Tfuel) Get() (cp models.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		reporterOne,
	}

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
	rep := ReporterOne{}
	req := request.
		New("GET", fmt.Sprintf(explorerOneURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Body.Balance.Tfuelwei, nil
}
