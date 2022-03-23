package eos

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
	address        = "bitfinexdep1"
	explorerOneURL = "https://eos.greymass.com/v1/chain/get_account"
	explorerTwoURL = "https://eos.greymass.com/v1/chain/get_currency_balance"
)

type Eos struct {
	Name    string
	Address string
}

func New() *Eos {
	return &Eos{
		Name:    "eos",
		Address: address,
	}
}

func (e *Eos) Get() (cp models.Checkpoint, err error) {
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
	var reqPld = struct {
		Code    string `json:"code"`
		Account string `json:"account"`
		Symbol  string `json:"symbol"`
	}{
		Code:    "eosio.token",
		Account: "bitfinexdep1",
		Symbol:  "EOS",
	}

	b, err := json.Marshal(reqPld)
	if err != nil {
		return decimal.NewFromInt(0), fmt.Errorf("eos reporter 2 %w", err)
	}

	rep := make(ExplorerTwoResponse, 0)
	req := request.
		New("POST", explorerTwoURL, bytes.NewReader(b)).
		AddHeaders("Content-Type", "application/json").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), fmt.Errorf("eos reporter 2 %w", err)
	}

	if len(rep) == 0 {
		return decimal.NewFromInt(0), errors.New("eos reporter 2 balance not found")
	}

	for _, v := range rep {
		if strings.HasSuffix(v, "EOS") {
			strBalPre := strings.Split(v, " ")
			strBalPost := strings.TrimSpace(strBalPre[0])
			return decimal.RequireFromString(strBalPost), nil
		}
	}

	return decimal.NewFromInt(0), errors.New("eos reporter 2 balance not found by token")
}
