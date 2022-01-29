package kardiachain

import (
	"fmt"
	"strings"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/gocolly/colly"
	"github.com/shopspring/decimal"
)

const (
	address        = "0x876EabF441B2EE5B5b0554Fd502a8E0600950cFa"
	explorerOneURL = "https://explorer.kardiachain.io/address/%s/transactions"
)

type Callisto struct {
	Name    string
	Address string
}

func New() *Callisto {
	return &Callisto{
		Name:    "kardiachain",
		Address: address,
	}
}

func (c *Callisto) Get() (cp models.Checkpoint, err error) {
	// only one explorer is known for kardiachain
	reporterPool := []reporterpool.Reporter{
		reporterOne,
	}

	for {
		var reporter reporterpool.Reporter
		reporterPool, reporter = reporterpool.PullRandReporter(reporterPool)

		balance, err := reporter(c.Address)
		if err == nil {
			cp.Balance = balance
			cp.Address = c.Address
			cp.Coin = c.Name
			break
		} else {
			// don't care about failure here since if 1 reporter fails, another might succeed
			fmt.Printf("failed getting %s wallet info: %s\n", c.Name, err)
		}

		if len(reporterPool) == 0 {
			break
		}
	}

	return
}

func reporterOne(address string) (decimal.Decimal, error) {
	var bal decimal.Decimal
	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("[data-test=address_balance]", func(e *colly.HTMLElement) {
		if len(e.Text) == 0 {
			return
		}

		split := strings.Split(e.Text, " ")
		if len(split) == 0 {
			return
		}

		str := split[0]
		if len(str) == 0 {
			return
		}

		str = strings.ReplaceAll(str, ",", "")
		str = strings.TrimSpace(str)
		bal = decimal.RequireFromString(str)
	})

	// Start scraping
	c.Visit(fmt.Sprintf(explorerOneURL, address))
	return bal, nil
}
