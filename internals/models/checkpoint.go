package models

import (
	"context"

	"github.com/cs-bird/internals/psql"
	"github.com/shopspring/decimal"
)

type Checkpoint struct {
	Height  int
	Coin    string
	Address string
	Balance decimal.Decimal
	Nonce   int64
}

func PersistCheckpoint(ctx context.Context, db psql.DBIterator, pld Checkpoint) (err error) {
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

	_, err = db.ExecContext(
		ctx,
		stmt,
		pld.Coin,
		pld.Address,
		pld.Balance.String(),
		pld.Nonce,
	)

	return
}
