package algorand

import (
	"errors"
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address         = "JDQ7EW3VY2ZHK4DKUHMNP35XLFPRJBND6M7SZ7W5RCFDNYAA47OC5IS62I"
	algoexplorerURL = "https://indexer.algoexplorerapi.io/v2/accounts/%s?include-all=true"
	purestakeURL    = "https://ms30v30mrj.execute-api.ca-central-1.amazonaws.com/Prod/algorand/mainnet/account/%s"
)

type Algorand struct {
	Name    string
	Address string
}

func New() *Algorand {
	return &Algorand{
		Name:    "algorand",
		Address: address,
	}
}

func (c *Algorand) Get() (cp types.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		algoexplorerReport,
		purestakeReport,
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

func algoexplorerReport(address string) (decimal.Decimal, error) {
	rep := AlgoexplorerReport{}
	req := request.
		New("GET", fmt.Sprintf(algoexplorerURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Account.Amount, nil
}

func purestakeReport(address string) (decimal.Decimal, error) {
	rep := PurestakeReport{}
	req := request.
		New("GET", fmt.Sprintf(purestakeURL, address), nil).
		AddHeaders("Host", "ms30v30mrj.execute-api.ca-central-1.amazonaws.com").
		AddHeaders("Origin", "https://goalseeker.purestake.io").
		AddHeaders("Referer", "https://goalseeker.purestake.io/algorand/mainnet/account/JDQ7EW3VY2ZHK4DKUHMNP35XLFPRJBND6M7SZ7W5RCFDNYAA47OC5IS62I").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Amount, nil
}
