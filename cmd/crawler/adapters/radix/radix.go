package radix

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/request"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	address        = "rdx1qspsgtz8wxzhs7h3m5dxz4uw0a350qp567zz4v9xgamcag8kavsxvpctcgge2"
	explorerOneURL = "https://mainnet.radixdlt.com/account/balances"
)

type Redix struct {
	Name    string
	Address string
}

func New() *Redix {
	return &Redix{
		Name:    "radix",
		Address: address,
	}
}

func (pd *Redix) Get() (cp models.Checkpoint, err error) {
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
		NetworkIdentifier struct {
			Network string `json:"network"`
		} `json:"network_identifier"`
		AccountIdentifier struct {
			Address string `json:"address"`
		} `json:"account_identifier"`
	}{
		NetworkIdentifier: struct {
			Network string `json:"network"`
		}{
			Network: "mainnet",
		},
		AccountIdentifier: struct {
			Address string `json:"address"`
		}{
			Address: "rdx1qspsgtz8wxzhs7h3m5dxz4uw0a350qp567zz4v9xgamcag8kavsxvpctcgge2",
		},
	}

	b, err := json.Marshal(reqPld)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	rep := RadixReport{}
	req := request.
		New("POST", explorerOneURL, bytes.NewReader(b)).
		AddHeaders("Content-Type", "application/json").
		AddHeaders("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36").
		AddHeaders("Referer", "https://explorer.radixdlt.com/").
		AddHeaders("Authority", "mainnet.radixdlt.com").
		AddHeaders("Origin", "https://explorer.radixdlt.com").
		AddHeaders("x-radixdlt-correlation-id", uuid.New().String()).
		AddHeaders("x-radixdlt-target-gw-api", "1.0.2").
		AddHeaders("x-radixdlt-method", "accountBalancesPost").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	for _, v := range rep.AccountBalances.LiquidBalances {
		if strings.HasPrefix(v.TokenIdentifier.Rri, "xrd_") {
			return v.Value, nil
		}
	}

	return decimal.NewFromInt(0), err
}
