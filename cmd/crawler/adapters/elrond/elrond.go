package elrond

import (
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address       = "erd1a56dkgcpwwx6grmcvw9w5vpf9zeq53w3w7n6dmxcpxjry3l7uh2s3h9dtr"
	elrondURL     = "https://api.elrond.com/accounts/%s"
	elrondscanURL = "https://api.elrondscan.com/accounts/%s"
)

type Elrond struct {
	Name    string
	Address string
}

func New() *Elrond {
	return &Elrond{
		Name:    "elrond",
		Address: address,
	}
}

func (e *Elrond) Get() (cp types.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		elrondReport,
		elrondscanReport,
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

func elrondReport(address string) (decimal.Decimal, error) {
	rep := ElrondReport{}
	req := request.
		New("GET", fmt.Sprintf(elrondURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Balance, nil
}

func elrondscanReport(address string) (decimal.Decimal, error) {
	rep := ElrondScanReport{}
	req := request.
		New("GET", fmt.Sprintf(elrondscanURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Balance, nil
}
