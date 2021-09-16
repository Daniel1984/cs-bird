package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cs-bird/cmd/crawler/adapters/cosmos"
	"github.com/cs-bird/cmd/crawler/adapters/dfinity"
	"github.com/cs-bird/cmd/crawler/adapters/elrond"
	"github.com/cs-bird/cmd/crawler/adapters/filecoin"
	"github.com/cs-bird/cmd/crawler/pipeline"
	"github.com/cs-bird/cmd/crawler/types"
	"github.com/cs-bird/internals/psql"
)

func main() {
	db, err := psql.New()
	if err != nil {
		log.Fatal(err)
	}

	pl := pipeline.New()
	pl.Add(filecoin.New())
	pl.Add(elrond.New())
	pl.Add(cosmos.New())
	pl.Add(dfinity.New())

	for {
		checkpoints := pl.Process()
		for coin, cp := range checkpoints {
			if cp.Err != nil {
				fmt.Printf("report to slack:> %s -> %+v\n", coin, cp)
				continue
			}

			if err := types.PersistCheckpoint(context.TODO(), db.Client, cp.Res); err != nil {
				fmt.Printf("failed persisting %s:>%+v, msg:%s\n", coin, cp, err)
			}
		}

		time.Sleep(120 * time.Second)
	}
}
