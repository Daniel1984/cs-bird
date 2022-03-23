package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/cs-bird/internals/psql"
)

type BalanceChange struct {
	BalanceChanges int    `json:"balance_changes"`
	Coin           string `json:"coin"`
	Address        string `json:"address"`
	Balance        string `json:"balance"`
}

type BalanceChangeRepo struct {
	DB psql.DBIterator
}

func (bcr BalanceChangeRepo) FetchForHours(ctx context.Context, hours int) ([]BalanceChange, error) {
	bcs := make([]BalanceChange, 0)

	query := []string{
		`
			SELECT
				COUNT(DISTINCT root.balance) - 1,
				root.coin,
				root.address,
				(
					SELECT
						balance
					FROM checkpoints
					WHERE address = root.address
					AND coin = root.coin
					ORDER BY time DESC
					LIMIT 1
				) as balace
		`,
		"FROM checkpoints AS root",
		fmt.Sprintf("WHERE root.time > NOW() - INTERVAL '%d hours'", hours),
		"GROUP BY root.coin, root.address",
	}

	rows, err := bcr.DB.QueryContext(ctx, strings.Join(query, " "))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var bc BalanceChange

		if err := rows.Scan(
			&bc.BalanceChanges,
			&bc.Coin,
			&bc.Address,
			&bc.Balance,
		); err != nil {
			return nil, err
		}

		bcs = append(bcs, bc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bcs, nil
}
