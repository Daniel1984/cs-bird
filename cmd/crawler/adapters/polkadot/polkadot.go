package polkadot

import (
	"bytes"
	"encoding/json"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "12T1tgaYZzEkFpnPvyqttmPRJxbGbR4uDx49cvZR5SRF8QDu"
	explorerOneURL = "https://polkadot.webapi.subscan.io/api/v2/scan/search"
)

type Polkadot struct {
	Name    string
	Address string
}

func New() *Polkadot {
	return &Polkadot{
		Name:    "polkadot",
		Address: address,
	}
}

func (pd *Polkadot) Get() (cp models.Checkpoint, err error) {
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
	var reqPld = struct {
		Key  string `json:"key"`
		Row  int    `json:"row"`
		Page int    `json:"page"`
	}{
		Key:  address,
		Row:  1,
		Page: 0,
	}

	b, err := json.Marshal(reqPld)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	rep := PolkadotReport{}
	req := request.
		New("POST", explorerOneURL, bytes.NewReader(b)).
		AddHeaders("Content-Type", "application/json").
		AddHeaders("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36").
		AddHeaders("Referer", "https://polkadot.subscan.io/").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Data.Account.Balance, nil
}
