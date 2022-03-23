package wax

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address        = "bitfinexwax1"
	explorerOneURL = "https://wax.greymass.com/v1/chain/get_account"
	explorerTwoURL = "https://wax.eosrio.io/v2/state/get_tokens?account=%s"
)

type Wax struct {
	Name    string
	Address string
}

func New() *Wax {
	return &Wax{
		Name:    "wax",
		Address: address,
	}
}

func (e *Wax) Get() (cp models.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{reporterOne, reporterTwo}

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
	var reqPld = struct {
		AccName string `json:"account_name"`
	}{
		AccName: address,
	}

	b, err := json.Marshal(reqPld)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	rep := ExplorerOneResponse{}
	req := request.
		New("POST", explorerOneURL, bytes.NewReader(b)).
		AddHeaders("Content-Type", "application/json").
		Do().
		Decode(&rep)

		// io.Copy(os.Stdout, req.Res.Body)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), fmt.Errorf("wax reporter 1 %w", err)
	}

	if len(rep.CoreLiquidBalance) == 0 {
		return decimal.NewFromInt(0), errors.New("wax reporter 1 could not get balance")
	}

	strBalPre := strings.Split(rep.CoreLiquidBalance, " ")
	strBalPost := strings.TrimSpace(strBalPre[0])
	return decimal.RequireFromString(strBalPost), nil
}

func reporterTwo(address string) (decimal.Decimal, error) {
	rep := ExplorerTwoResponse{}
	req := request.
		New("GET", fmt.Sprintf(explorerTwoURL, address), nil).
		AddHeaders("Content-Type", "application/json").
		AddHeaders("Referer", "https://wax.bloks.io/").
		AddHeaders("Origin", "https://wax.bloks.io").
		Do().
		Decode(&rep)

	// io.Copy(os.Stdout, req.Res.Body)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	for _, v := range rep.Tokens {
		if v.Symbol == "WAX" {
			return v.Amount, nil
		}
	}

	return decimal.NewFromInt(0), errors.New("wax reporter 2 could not get balance")
}
