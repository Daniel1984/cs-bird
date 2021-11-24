package tezos

import (
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "tz1KtGwriE7VuLwT3LwuvU9Nv4wAxP7XZ57d"
	explorerOneURL = "https://api.tzkt.io/v1/accounts/%s?metadata=true"
	explorerTwoURL = "https://api.tzstats.com/explorer/account/%s?meta=1"
)

type Tezos struct {
	Name    string
	Address string
}

func New() *Tezos {
	return &Tezos{
		Name:    "tezos",
		Address: address,
	}
}

func (e *Tezos) Get() (cp types.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		reporterOne,
		reporterTwo,
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

	return rep.Balance.Div(decimal.NewFromInt(1000000)).Round(6), nil
}

func reporterTwo(address string) (decimal.Decimal, error) {
	rep := ReporterTwo{}
	req := request.
		New("GET", fmt.Sprintf(explorerTwoURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Balance.Round(6), nil
}
