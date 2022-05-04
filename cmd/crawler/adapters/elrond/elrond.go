package elrond

import (
	"fmt"

	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/reporterpool"
	"github.com/cs-bird/internals/request"
	"github.com/shopspring/decimal"
)

const (
	address   = "erd1a56dkgcpwwx6grmcvw9w5vpf9zeq53w3w7n6dmxcpxjry3l7uh2s3h9dtr"
	explorer1 = "https://api.elrond.com/accounts/%s"
	explorer2 = "https://elrondscan.com:3000/accounts/%s"
	explorer3 = "https://internal-api.elrond.com/accounts/%s"
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

func (e *Elrond) Get() (cp models.Checkpoint, err error) {
	reporterPool := []reporterpool.Reporter{
		explorer1Report,
		explorer2Report,
		explorer3Report,
	}

	for {
		var reporter reporterpool.Reporter
		reporterPool, reporter = reporterpool.PullRandReporter(reporterPool)

		balance, err := reporter(e.Address)
		if err == nil {
			cp.Balance = balance
			cp.Address = e.Address
			cp.Coin = e.Name
			break
		} else {
			// don't care about failure here since if 1 reporter fails, another might succeed
			fmt.Printf("failed getting %s wallet info: %s\n", e.Name, err)
		}

		if len(reporterPool) == 0 {
			break
		}
	}

	return
}

func explorer1Report(address string) (decimal.Decimal, error) {
	rep := Explorer1Report{}
	req := request.
		New("GET", fmt.Sprintf(explorer1, address), nil).
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Balance, nil
}

func explorer2Report(address string) (decimal.Decimal, error) {
	rep := Explorer2Report{}
	req := request.
		New("GET", fmt.Sprintf(explorer2, address), nil).
		AddHeaders("Accept", "application/json").
		AddHeaders("Accept-Language", "en-GB,en;q=0.9,lt-GB;q=0.8,lt;q=0.7,en-US;q=0.6").
		AddHeaders("Authorization", "WlhKa01YRnhjWEZ4Y1hGeGNYRnhjWEZ4Y1hCeGNYRnhjWEZ4Y1hGeGNYRnhjWEZ4Y1hGeGNYRnhjWEZ4Y1hGeGNYRnhlV3hzYkhOc2JYRTJlVFk2TWpBeU1pMHdOUzB3TkZReE13PT0=").
		AddHeaders("Connection", "keep-alive").
		AddHeaders("Origin", "https://elrondscan.com").
		AddHeaders("Referer", "https://elrondscan.com/").
		AddHeaders("Sec-Fetch-Dest", "empty").
		AddHeaders("Sec-Fetch-Mode", "cors").
		AddHeaders("Sec-Fetch-Site", "same-site").
		AddHeaders("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36").
		AddHeaders("X-Requested-With", "XMLHttpRequest").
		AddHeaders("X-Socket-Id", "319760622.309542742").
		AddHeaders("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100""`).
		AddHeaders("sec-ch-ua-mobile", "?0").
		AddHeaders("sec-ch-ua-platform", "macOS").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Balance, nil
}

func explorer3Report(address string) (decimal.Decimal, error) {
	rep := Explorer3Report{}
	req := request.
		New("GET", fmt.Sprintf(explorer3, address), nil).
		AddHeaders("authority", "internal-api.elrond.com").
		AddHeaders("accept", "*/*").
		AddHeaders("accept-language", "en-GB,en;q=0.9,lt-GB;q=0.8,lt;q=0.7,en-US;q=0.6").
		AddHeaders("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NBZGRyZXNzIjoiZXJkMWs0ZHVjMm5leWRod2FucW5ua2ZlN3p2M2QzdG1qMDAzd3Zlc3dsYXBmc3JsbXc0Zjc3eHNqeGFxMmMiLCJzaWduYXR1cmUiOiI2NDVhNTg0Y2QzMjczZDg5NmIyYzQ5MzUyN2U3ZWYwZTJlNDA1NjM2NjhjNWZlZmU3MjZkYzFkMWIwZjFlMGNlNDBkZWQ2M2FiMzc3YzUzY2JhMDJhNTFjNzYxM2RiZDliOGNlZGYxNzNlYWMwNjdmNGFlYTZkYWYwMWI3MDcwMCIsImlhdCI6MTY1MTY3MDY0NCwiZXhwIjoxNjUxNjc3ODQ0fQ.TRfhtEf8Rxe67hgmvsWiyj-mkGA3HG4vv4BWmoBEdfU").
		AddHeaders("origin", "https://explorer.elrond.com").
		AddHeaders("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="100", "Google Chrome";v="100""`).
		AddHeaders("sec-ch-ua-mobile", "?0").
		AddHeaders("sec-ch-ua-platform", "macOS").
		AddHeaders("sec-fetch-dest", "empty").
		AddHeaders("sec-fetch-mode", "cors").
		AddHeaders("sec-fetch-site", "same-site").
		AddHeaders("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36").
		Do().
		Decode(&rep)

	if err := req.HasError(); err != nil {
		return decimal.NewFromInt(0), err
	}

	return rep.Balance, nil
}
