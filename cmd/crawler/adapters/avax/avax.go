package avax

import (
	"errors"
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address = "avax1ju2khnvkex4u30rt5jm5x8jd3jdfhvawg7ed9x"
	avaxURL = "https://explorerapi.avax.network/v2/addresses?address=%s&chainID=2oYMBNV4eNHyqk2fjjV5nVQLDbtmNJzq5s3qs3Lo6ftnC6FByM"
)

type Avax struct {
	Name    string
	Address string
}

func New() *Avax {
	return &Avax{
		Name:    "avax",
		Address: address,
	}
}

func (e *Avax) Get() (cp models.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{avaxReport}

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

func avaxReport(address string) (decimal.Decimal, error) {
	rep := AvaxReport{}
	req := request.
		New("GET", fmt.Sprintf(avaxURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	if len(rep.Addresses) == 0 {
		return decimal.NewFromInt(0), errors.New("no addresses returned for avax")
	}

	for _, v := range rep.Addresses {
		if v.Address == address {
			return v.Assets.Info.Balance, nil
		}
	}

	return decimal.NewFromInt(0), errors.New("avax hw address not found")
}
