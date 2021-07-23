package dfinity

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/request"
)

const (
	address = "ed753e2282b90f0c92f2ec2b21161b05c505f1c2a606dbe55d5d27872623395b"
	infoURL = "https://rosetta-api.internetcomputer.org/account/balance"
)

type Dfinity struct {
	Name    string
	Address string
}

func New() *Dfinity {
	return &Dfinity{
		Name:    "dfinity",
		Address: address,
	}
}

func (this *Dfinity) Get() (cp types.Checkpoint, err error) {
	balRep, err := getBalanceReport(this.Address)
	if err != nil {
		return cp, err
	}

	if len(balRep.Balances) == 0 {
		return cp, fmt.Errorf("unable to get dfinity balance")
	}

	cp.Balance = balRep.Balances[0].Value
	cp.Address = this.Address
	cp.Coin = this.Name

	return
}

func getBalanceReport(address string) (BalanceReport, error) {
	var reqPld = struct {
		NetworkIdentifier NetworkIdentifier `json:"network_identifier"`
		AccountIdentifier AccountIdentifier `json:"account_identifier"`
	}{
		NetworkIdentifier: NetworkIdentifier{
			Blockchain: "Internet Computer",
			Network:    "00000000000000020101",
		},
		AccountIdentifier: AccountIdentifier{
			Address: address,
		},
	}

	b, err := json.Marshal(reqPld)
	if err != nil {
		return BalanceReport{}, fmt.Errorf("marshal:> %s", err)
	}

	rep := BalanceReport{}
	req := request.
		New("POST", infoURL, bytes.NewReader(b)).
		AddHeaders("Content-Type", "application/json").
		Do().
		Decode(&rep)

	return rep, req.HasError()
}
