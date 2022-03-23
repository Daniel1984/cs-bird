package certic

import (
	"errors"
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "certik1azxdh422fnx4df78xsk3j4vsz2wl224wt5tfwd"
	explorerOneURL = "https://azuredragon.noopsbycertik.com/cosmos/bank/v1beta1/balances/%s"
)

type Certic struct {
	Name    string
	Address string
}

func New() *Certic {
	return &Certic{
		Name:    "certic",
		Address: address,
	}
}

func (e *Certic) Get() (cp models.Checkpoint, err error) {
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
		AddHeaders("Origin", "https://explorer.shentu.technology").
		AddHeaders("Referer", "https://explorer.shentu.technology/").
		AddHeaders("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	for _, v := range rep.Balances {
		if v.Denom == "uctk" {
			return v.Amount.Div(decimal.NewFromInt(1000000)).Round(6), nil
		}
	}

	return decimal.NewFromInt(0), errors.New("certic explorer 1 unable to get balance")
}
