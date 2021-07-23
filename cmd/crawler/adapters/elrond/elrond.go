package elrond

import (
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/request"
)

const (
	address = "erd1a56dkgcpwwx6grmcvw9w5vpf9zeq53w3w7n6dmxcpxjry3l7uh2s3h9dtr"
	infoURL = "https://api.elrond.com/accounts/%s"
)

type Elrond struct {
	Name    string
	Address string
}

func New() *Elrond {
	return &Elrond{
		Name:    "elrond",
		Address: address,
	}
}

func (this *Elrond) Get() (cp types.Checkpoint, err error) {
	rep, err := getReport(this.Address)
	if err != nil {
		return cp, err
	}

	cp.Balance = rep.Balance
	cp.Height = rep.Nonce
	cp.Address = this.Address
	cp.Coin = this.Name

	return
}

func getReport(address string) (Report, error) {
	rep := Report{}
	req := request.
		New("GET", fmt.Sprintf(infoURL, address), nil).
		Do().
		Decode(&rep)

	return rep, req.HasError()
}
