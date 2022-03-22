package omni

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "1BzjoD8rMvCjTDHhc7mUzG3ZB7idTjfW6U"
	explorerOneURL = "https://api.omniexplorer.info/v1/address/addr/"
)

type Omni struct {
	Name    string
	Address string
}

func New() *Omni {
	return &Omni{
		Name:    "omni",
		Address: address,
	}
}

func (pd *Omni) Get() (cp models.Checkpoint, err error) {
	repOne, err := getBalanceReport(pd.Address)
	if err != nil {
		return cp, err
	}

	cp.Balance = repOne
	cp.Address = pd.Address
	cp.Coin = pd.Name

	return
}

func getBalanceReport(address string) (decimal.Decimal, error) {
	rep := OmniReporterOne{}
	req := request.
		New("POST", explorerOneURL, strings.NewReader(fmt.Sprintf("addr=%s", address))).
		AddHeaders("Content-Type", "application/x-www-form-urlencoded").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	for _, v := range rep.Balance {
		if v.Symbol == "OMNI" {
			return v.Value, nil
		}
	}

	return decimal.NewFromInt(0), errors.New("omni value not found")
}
