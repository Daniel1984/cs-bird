package polygon

import (
	"context"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

const (
	address    = "0xf63724c70bf94b7ceefa0bc8e2de5948eded0800"
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

func (p *Polygon) Get() (cp types.Checkpoint, err error) {
	client, err := ethclient.Dial(polygonURL)
	if err != nil {
		return types.Checkpoint{}, err
	}

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.TODO(), account, nil)
	if err != nil {
		return types.Checkpoint{}, err
	}

	cp.Balance = decimal.NewFromBigInt(balance, 0)
	cp.Address = p.Address
	cp.Coin = p.Name
	return
}
