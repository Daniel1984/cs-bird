package cosmos

import (
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/request"
)

const (
	address = "cosmos1h9ymfm2fxrqgd257dlw5nku3jgqjgpl59sm5ns"
	infoURL = "https://lcd-cosmos.cosmostation.io/cosmos/bank/v1beta1/balances/%s"
)

type Cosmos struct {
	Name    string
	Address string
}

func New() *Cosmos {
	return &Cosmos{
		Name:    "cosmos",
		Address: address,
	}
}

func (this *Cosmos) Get() (cp types.Checkpoint, err error) {
	balRep, err := getBalanceReport(this.Address)
	if err != nil {
		return cp, err
	}

	if len(balRep.Balances) == 0 {
		return cp, fmt.Errorf("unable to get cosmos balance")
	}

	cp.Balance = balRep.Balances[0].Amount
	cp.Address = this.Address
	cp.Coin = this.Name

	return
}

func getBalanceReport(address string) (BalanceReport, error) {
	rep := BalanceReport{}
	req := request.
		New("GET", fmt.Sprintf(infoURL, address), nil).
		Do().
		Decode(&rep)

	return rep, req.HasError()
}
