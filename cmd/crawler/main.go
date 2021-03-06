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
	"github.com/cs-bird/cmd/crawler/adapters/certic"
	"github.com/cs-bird/cmd/crawler/adapters/cosmos"
	"github.com/cs-bird/cmd/crawler/adapters/dfinity"
	"github.com/cs-bird/cmd/crawler/adapters/elrond"
	"github.com/cs-bird/cmd/crawler/adapters/eos"
	"github.com/cs-bird/cmd/crawler/adapters/filecoin"
	"github.com/cs-bird/cmd/crawler/adapters/kardiachain"
	"github.com/cs-bird/cmd/crawler/adapters/kusama"
	"github.com/cs-bird/cmd/crawler/adapters/luna"
	"github.com/cs-bird/cmd/crawler/adapters/near"
	"github.com/cs-bird/cmd/crawler/adapters/omni"
	"github.com/cs-bird/cmd/crawler/adapters/polkadot"
	"github.com/cs-bird/cmd/crawler/adapters/polygon"
	"github.com/cs-bird/cmd/crawler/adapters/radix"
	"github.com/cs-bird/cmd/crawler/adapters/rose"
	"github.com/cs-bird/cmd/crawler/adapters/serum"
	"github.com/cs-bird/cmd/crawler/adapters/solana"
	"github.com/cs-bird/cmd/crawler/adapters/songbird"
	"github.com/cs-bird/cmd/crawler/adapters/stellar"
	"github.com/cs-bird/cmd/crawler/adapters/tezos"
	"github.com/cs-bird/cmd/crawler/adapters/tfuel"
	"github.com/cs-bird/cmd/crawler/adapters/theta"
	"github.com/cs-bird/cmd/crawler/adapters/tron"
	"github.com/cs-bird/cmd/crawler/adapters/ust"
	"github.com/cs-bird/cmd/crawler/adapters/vsys"
	"github.com/cs-bird/cmd/crawler/adapters/waves"
	"github.com/cs-bird/cmd/crawler/adapters/wax"
	"github.com/cs-bird/cmd/crawler/adapters/xinfin"
	"github.com/cs-bird/cmd/crawler/adapters/zil"
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
	pl.Add(songbird.New())
	pl.Add(near.New())
	pl.Add(omni.New())
	pl.Add(polkadot.New())
	pl.Add(kusama.New())
	pl.Add(radix.New())
	pl.Add(solana.New())
	pl.Add(stellar.New())
	pl.Add(tron.New())
	pl.Add(vsys.New())
	pl.Add(wax.New())
	pl.Add(waves.New())
	pl.Add(xinfin.New())
	pl.Add(zil.New())
	pl.Add(certic.New())
	pl.Add(eos.New())
	pl.Add(serum.New())
	pl.Add(luna.New())
	pl.Add(ust.New())

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

		interval := randnum.InRange(120, 600)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
