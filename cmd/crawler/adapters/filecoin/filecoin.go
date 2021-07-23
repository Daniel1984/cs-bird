package filecoin

import (
	"fmt"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/request"
)

const (
	address    = "f3rgsvwupbysklbwdx2444xximkapartwtue3aoyzws7z4cv23ku2nzcr32ioyrnjm7wulytyo36fgratyri5a"
	msgsURL    = "https://filfox.info/api/v1/address/%s/messages?pageSize=20&page=0"
	balanceURL = "https://filfox.info/api/v1/address/%s/balance-stats?duration=24h&samples=1"
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

func (this *Filecoin) Get() (cp types.Checkpoint, err error) {
	msgs, err := getMessages(this.Address)
	if err != nil {
		return cp, err
	}

	balRep, err := getBalance(this.Address)
	if err != nil {
		return cp, err
	}

	if len(balRep) == 0 {
		return cp, fmt.Errorf("missing balance report")
	}

	cp.Balance = balRep[0].Balance
	cp.Height = msgs.TotalCount
	cp.Address = this.Address
	cp.Coin = this.Name

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

func getBalance(address string) (BalanceReport, error) {
	br := BalanceReport{}
	req := request.
		New("GET", fmt.Sprintf(balanceURL, address), nil).
		Do().
		Decode(&br)

	return br, req.HasError()
}
