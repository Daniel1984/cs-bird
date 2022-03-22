package near

import "github.com/shopspring/decimal"

type ExplorerOneResponse struct {
	PageProps struct {
		Account struct {
			AccountID                string          `json:"accountId"`
			CreatedByTransactionHash string          `json:"createdByTransactionHash"`
			CreatedAtBlockTimestamp  int64           `json:"createdAtBlockTimestamp"`
			StorageUsage             string          `json:"storageUsage"`
			StakedBalance            decimal.Decimal `json:"stakedBalance"`
			NonStakedBalance         decimal.Decimal `json:"nonStakedBalance"`
			MinimumBalance           decimal.Decimal `json:"minimumBalance"`
			AvailableBalance         decimal.Decimal `json:"availableBalance"`
			TotalBalance             decimal.Decimal `json:"totalBalance"`
		} `json:"account"`
	} `json:"pageProps"`
}
