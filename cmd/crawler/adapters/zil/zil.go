package zil

import (
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "zil1xfsrre5qgx0mqg99xc0l2cuyu9ntt259ngsu7s"
	explorerOneURL = "https://api.viewblock.io/zilliqa/addresses/%s?network=mainnet&page=1"
)

type Zil struct {
	Name    string
	Address string
}

func New() *Zil {
	return &Zil{
		Name:    "zil",
		Address: address,
	}
}

func (e *Zil) Get() (cp models.Checkpoint, err error) {
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
		AddHeaders("Content-Type", "application/json").
		AddHeaders("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36").
		AddHeaders("Origin", "https://viewblock.io").
		AddHeaders("Referer", "https://viewblock.io/").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Balance.Div(decimal.NewFromInt(1000000000000)).Round(12), nil
}
