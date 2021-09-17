package filecoin

import (
	"errors"
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address     = "f3rgsvwupbysklbwdx2444xximkapartwtue3aoyzws7z4cv23ku2nzcr32ioyrnjm7wulytyo36fgratyri5a"
	msgsURL     = "https://filfox.info/api/v1/address/%s/messages?pageSize=20&page=0"
	filfoxURL   = "https://filfox.info/api/v1/address/%s/balance-stats?duration=24h&samples=48"
	filscoutURL = "https://api.filscout.com/api/v1/actor/balance/day/%s"
)

type Filecoin struct {
	Name    string
	Address string
}

func New() *Filecoin {
	return &Filecoin{
		Name:    "filecoin",
		Address: address,
	}
}

func (f *Filecoin) Get() (cp types.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		filfoxReport,
		filscoutReport,
	}

	for {
		var reporter reporterpool.Reporter
		reporterPool, reporter = reporterpool.PullRandReporter(reporterPool)

		balance, err := reporter(f.Address)
		if err == nil {
			cp.Balance = balance
			cp.Address = f.Address
			cp.Coin = f.Name
			break
		} else {
			// don't care about failure here since if 1 reporter fails, another might succeed
			fmt.Printf("failed getting %s wallet info: %s\n", f.Name, err)
		}

		if len(reporterPool) == 0 {
			err = errors.New("run out of connections for: " + f.Name)
			break
		}
	}

	return
}

func getMessages(address string) (MessagesReport, error) {
	mr := MessagesReport{}
	req := request.
		New("GET", fmt.Sprintf(msgsURL, address), nil).
		Do().
		Decode(&mr)

	return mr, req.HasError()
}

func filfoxReport(address string) (decimal.Decimal, error) {
	rep := FilfoxReport{}
	req := request.
		New("GET", fmt.Sprintf(filfoxURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	if len(rep) == 0 {
		return decimal.NewFromInt(0), fmt.Errorf("filfox: missing balances")
	}

	latestBalance := rep[len(rep)-1].Balance
	formatedBalance := latestBalance.Div(decimal.NewFromInt(1000000000000000000)).Round(6)
	return formatedBalance, nil
}

func filscoutReport(address string) (decimal.Decimal, error) {
	rep := FilscoutReport{}
	req := request.
		New("GET", fmt.Sprintf(filscoutURL, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	if len(rep.Data.Balance) == 0 {
		return decimal.NewFromInt(0), fmt.Errorf("filscout: missing balances")
	}

	return rep.Data.Balance[len(rep.Data.Balance)-1], nil
}
