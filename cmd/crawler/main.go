package main

import (
	"fmt"
	"time"

	"github.com/cs-bird/cmd/crawler/adapters/wax"
	"github.com/cs-bird/cmd/crawler/pipeline"
	"github.com/cs-bird/internals/randnum"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// db, err := psql.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// repo := models.NewRepository(db.Client)

	pl := pipeline.New()
	// pl.Add(rose.New())
	// pl.Add(filecoin.New())
	// pl.Add(elrond.New())
	// pl.Add(cosmos.New())
	// pl.Add(dfinity.New())
	// pl.Add(algorand.New())
	// pl.Add(aion.New())
	// pl.Add(avax.New())
	// pl.Add(polygon.New())
	// pl.Add(callisto.New())
	// pl.Add(tezos.New())
	// pl.Add(theta.New())
	// pl.Add(tfuel.New())
	// pl.Add(kardiachain.New())
	// pl.Add(songbird.New())
	// pl.Add(near.New())
	// pl.Add(omni.New())
	// pl.Add(polkadot.New())
	// pl.Add(kusama.New())
	// pl.Add(radix.New())
	// pl.Add(solana.New())
	// pl.Add(stellar.New())
	// pl.Add(tron.New())
	// pl.Add(vsys.New())
	pl.Add(wax.New())

	for {
		checkpoints := pl.Process()
		for coin, cp := range checkpoints {
			if cp.Err != nil {
				fmt.Printf("report to slack:> %s -> %+v\n", coin, cp)
				continue
			}

			// if len(cp.Res.Address) > 0 && len(cp.Res.Coin) > 0 {
			// 	if err := repo.Checkpoint.Insert(context.TODO(), cp.Res); err != nil {
			// 		fmt.Printf("failed persisting %s:>%+v, msg:%s\n", coin, cp, err)
			// 	}
			// }
			fmt.Printf("----------> %+v\n", cp.Res)
		}

		interval := randnum.InRange(300, 600)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
