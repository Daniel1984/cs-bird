package polygon

import (
	"context"

	"github.com/cs-bird/internals/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

const (
	// address    = "0xf63724c70bf94b7ceefa0bc8e2de5948eded0800"
	address    = "0x876EabF441B2EE5B5b0554Fd502a8E0600950cFa"
	polygonURL = "https://rpc-mainnet.maticvigil.com/"
)

type Polygon struct {
	Name    string
	Address string
}

func New() *Polygon {
	return &Polygon{
		Name:    "polygon",
		Address: address,
	}
}

func (p *Polygon) Get() (cp models.Checkpoint, err error) {
	client, err := ethclient.Dial(polygonURL)
	if err != nil {
		return models.Checkpoint{}, err
	}

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.TODO(), account, nil)
	if err != nil {
		return models.Checkpoint{}, err
	}

	cp.Balance = decimal.NewFromBigInt(balance, 0)
	cp.Address = p.Address
	cp.Coin = p.Name
	return
}
