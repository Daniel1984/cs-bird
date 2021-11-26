package models

import (
	"context"

	"github.com/cs-bird/internals/psql"
	"github.com/shopspring/decimal"
)

type Checkpoint struct {
	Height  int
	Nonce   int64
	Balance decimal.Decimal
	Coin    string
	Address string
}

type CheckpointRepo struct {
	DB psql.DBIterator
}

func (cpr CheckpointRepo) Insert(ctx context.Context, pld Checkpoint) (err error) {
	stmt := `
		INSERT INTO
		public."checkpoints" (
			time,
			coin,
			address,
			balance,
			nonce
		)
		VALUES (NOW(), $1, $2, $3, $4)
	`

	_, err = cpr.DB.ExecContext(
		ctx,
		stmt,
		pld.Coin,
		pld.Address,
		pld.Balance.String(),
		pld.Nonce,
	)

	return
}
