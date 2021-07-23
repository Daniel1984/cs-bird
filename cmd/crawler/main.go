package main

import (
	"fmt"
	"time"

	"github.com/cs-bird/cmd/crawler/adapters/cosmos"
	"github.com/cs-bird/cmd/crawler/adapters/dfinity"
	"github.com/cs-bird/cmd/crawler/adapters/elrond"
	"github.com/cs-bird/cmd/crawler/adapters/filecoin"
	"github.com/cs-bird/cmd/crawler/pipeline"
)

func main() {
	pl := pipeline.New()
	pl.Add(filecoin.New())
	pl.Add(elrond.New())
	pl.Add(cosmos.New())
	pl.Add(dfinity.New())

	for {
		report := pl.Process()
		fmt.Printf("pipeline response: %+v\n", report)
		time.Sleep(60 * time.Second)
	}
}
