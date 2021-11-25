package cosmos

import (
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address          = "cosmos1h9ymfm2fxrqgd257dlw5nku3jgqjgpl59sm5ns"
	cosmosstationURL = "https://lcd-cosmos.cosmostation.io/cosmos/bank/v1beta1/balances/%s"
	atomscanURL      = "https://node.atomscan.com/bank/balances/%s"
)

type Cosmos struct {
	Name    string
	Address string
}

func New() *Cosmos {
	return &Cosmos{
		Name:    "cosmos",
		Address: address,
	}
}

func (c *Cosmos) Get() (cp models.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		cosmosstationReport,
		atomscanReport,
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
			break
		}
	}

	return
}

func cosmosstationReport(address string) (decimal.Decimal, error) {
	rep := CosmosstationReport{}
	req := request.
		New("GET", fmt.Sprintf(cosmosstationURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	if len(rep.Balances) == 0 {
		return decimal.NewFromInt(0), fmt.Errorf("unable to get balance from cosmosstation")
	}

	return rep.Balances[0].Amount, nil
}

func atomscanReport(address string) (decimal.Decimal, error) {
	rep := AtomscanReport{}
	req := request.
		New("GET", fmt.Sprintf(atomscanURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	if len(rep.Result) == 0 {
		return decimal.NewFromInt(0), fmt.Errorf("unable to get balance from atomscan")
	}

	return rep.Result[0].Amount, nil
}
