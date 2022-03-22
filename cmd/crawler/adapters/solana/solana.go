package solana

import (
	"bytes"
	"encoding/json"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/request"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	address        = "FxteHmLwG9nk1eL4pjNve3Eub2goGkkz6g6TbvdmW46a"
	explorerOneURL = "https://explorer-api.mainnet-beta.solana.com"
)

type Solana struct {
	Name    string
	Address string
}

func New() *Solana {
	return &Solana{
		Name:    "solana",
		Address: address,
	}
}

func (pd *Solana) Get() (cp models.Checkpoint, err error) {
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
		Method  string        `json:"method"`
		Jsonrpc string        `json:"jsonrpc"`
		Params  []interface{} `json:"params"`
		ID      string        `json:"id"`
	}{
		Method:  "getAccountInfo",
		Jsonrpc: "2.0",
		ID:      uuid.New().String(),
		Params: []interface{}{
			"FxteHmLwG9nk1eL4pjNve3Eub2goGkkz6g6TbvdmW46a",
			map[string]string{
				"encoding":   "jsonParsed",
				"commitment": "confirmed",
			},
		},
	}

	b, err := json.Marshal(reqPld)
	if err != nil {
		return decimal.NewFromInt(0), err
	}

	rep := SolanaReport{}
	req := request.
		New("POST", explorerOneURL, bytes.NewReader(b)).
		AddHeaders("Content-Type", "application/json").
		AddHeaders("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36").
		AddHeaders("Referer", "https://explorer.solana.com/").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Result.Value.Lamports, nil
}
