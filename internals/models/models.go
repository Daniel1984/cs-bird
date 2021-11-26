package models

import (
	"context"

	"github.com/cs-bird/internals/psql"
)

type Repo struct {
	Checkpoint interface {
		Insert(ctx context.Context, cp Checkpoint) error
	}
	BalanceChange interface {
		FetchForHours(ctx context.Context, hours int) ([]BalanceChange, error)
	}
}

func NewRepository(db psql.DBIterator) Repo {
	return Repo{
		BalanceChange: BalanceChangeRepo{DB: db},
		Checkpoint:    CheckpointRepo{DB: db},
	}
}
