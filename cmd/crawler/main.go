package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cs-bird/cmd/crawler/adapters/aion"
	"github.com/cs-bird/cmd/crawler/adapters/algorand"
	"github.com/cs-bird/cmd/crawler/adapters/avax"
	"github.com/cs-bird/cmd/crawler/adapters/callisto"
	"github.com/cs-bird/cmd/crawler/adapters/cosmos"
	"github.com/cs-bird/cmd/crawler/adapters/dfinity"
	"github.com/cs-bird/cmd/crawler/adapters/elrond"
	"github.com/cs-bird/cmd/crawler/adapters/filecoin"
	"github.com/cs-bird/cmd/crawler/adapters/kardiachain"
	"github.com/cs-bird/cmd/crawler/adapters/polygon"
	"github.com/cs-bird/cmd/crawler/adapters/rose"
	"github.com/cs-bird/cmd/crawler/adapters/tezos"
	"github.com/cs-bird/cmd/crawler/adapters/tfuel"
	"github.com/cs-bird/cmd/crawler/adapters/theta"
	"github.com/cs-bird/cmd/crawler/pipeline"
	"github.com/cs-bird/internals/models"
	"github.com/cs-bird/internals/psql"
	"github.com/cs-bird/internals/randnum"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := psql.New()
	if err != nil {
		log.Fatal(err)
	}

	repo := models.NewRepository(db.Client)

	pl := pipeline.New()
	pl.Add(rose.New())
	pl.Add(filecoin.New())
	pl.Add(elrond.New())
	pl.Add(cosmos.New())
	pl.Add(dfinity.New())
	pl.Add(algorand.New())
	pl.Add(aion.New())
	pl.Add(avax.New())
	pl.Add(polygon.New())
	pl.Add(callisto.New())
	pl.Add(tezos.New())
	pl.Add(theta.New())
	pl.Add(tfuel.New())
	pl.Add(kardiachain.New())

	for {
		checkpoints := pl.Process()
		for coin, cp := range checkpoints {
			if cp.Err != nil {
				fmt.Printf("report to slack:> %s -> %+v\n", coin, cp)
				continue
			}

			if len(cp.Res.Address) > 0 && len(cp.Res.Coin) > 0 {
				if err := repo.Checkpoint.Insert(context.TODO(), cp.Res); err != nil {
					fmt.Printf("failed persisting %s:>%+v, msg:%s\n", coin, cp, err)
				}
			}
		}

		interval := randnum.InRange(60, 90)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
