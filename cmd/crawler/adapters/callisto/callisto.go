package callisto

import (
	"fmt"
	"strings"

	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/gocolly/colly"
	"github.com/shopspring/decimal"
)

const (
	address     = "0x16c633f1AFC23DC14E29e2c3220a4b5E4a5260a0"
	explorerOne = "https://explorer.callisto.network/address/%s"
	explorerTwo = "https://callistoexplorer.com/address/%s"
)

type Callisto struct {
	Name    string
	Address string
}

func New() *Callisto {
	return &Callisto{
		Name:    "callisto",
		Address: address,
	}
}

func (c *Callisto) Get() (cp types.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		reporterOne,
		reporterTwo,
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

	c.OnHTML("h3[data-test=address_balance]", func(e *colly.HTMLElement) {
		if len(e.Text) > 0 && strings.HasSuffix(e.Text, "CLO") {
			str := strings.ReplaceAll(e.Text, ",", "")
			str = strings.ReplaceAll(str, "CLO", "")
			str = strings.TrimSpace(str)
			bal = decimal.RequireFromString(str)
		}
	})

	// Start scraping
	c.Visit(fmt.Sprintf(explorerOne, address))

	return bal, nil
}

func reporterTwo(address string) (decimal.Decimal, error) {
	var bal decimal.Decimal
	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML(".address-general-info .content ul li", func(e *colly.HTMLElement) {
		if len(e.Text) > 0 && strings.HasPrefix(e.Text, "Balance") && strings.HasSuffix(e.Text, "CLO") {
			str := strings.ReplaceAll(e.Text, "Balance", "")
			str = strings.ReplaceAll(str, "CLO", "")
			str = strings.TrimSpace(str)
			bal = decimal.RequireFromString(str)
		}
	})

	// Start scraping
	c.Visit(fmt.Sprintf(explorerTwo, address))

	return bal, nil
}
