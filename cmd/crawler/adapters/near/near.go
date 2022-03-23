package near

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	address        = "383e50ea1a754ed3acd0d59116f221add87adb82559f31ca6d377f058fe83375"
	explorerOneURL = "https://explorer.near.org/_next/data/qmZoGxGkMm0Xv511MFKII/accounts/%s.json?id=%s"
	explorerTwoURL = "http://rpc.mainnet.near.org/"
)

type Near struct {
	Name    string
	Address string
}

func New() *Near {
	return &Near{
		Name:    "near",
		Address: address,
	}
}

func (e *Near) Get() (cp models.Checkpoint, err error) {
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
	var reqPld = struct {
		Jsonrpc string `json:"jsonrpc"`
		ID      string `json:"id"`
		Method  string `json:"method"`
		Params  struct {
			RequestType string `json:"request_type"`
			Finality    string `json:"finality"`
			AccountID   string `json:"account_id"`
		} `json:"params"`
	}{
		Jsonrpc: "2.0",
		ID:      uuid.New().String(),
		Method:  "query",
		Params: struct {
			RequestType string `json:"request_type"`
			Finality    string `json:"finality"`
			AccountID   string `json:"account_id"`
		}{
			RequestType: "view_account",
			Finality:    "final",
			AccountID:   address,
		},
	}

	b, err := json.Marshal(reqPld)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	rep := ExplorerOneResponse{}
	req := request.
		New("POST", explorerTwoURL, bytes.NewReader(b)).
		AddHeaders("Content-Type", "application/json").
		Do().
		Decode(&rep)
	io.Copy(os.Stdout, req.Res.Body)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Result.Amount, nil
}
