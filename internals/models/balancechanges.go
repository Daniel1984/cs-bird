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
	Address        string `joisn:"address"`
}

type BalanceChangeRepo struct {
	DB psql.DBIterator
}

func (bcr BalanceChangeRepo) FetchForHours(ctx context.Context, hours int) ([]BalanceChange, error) {
	bcs := make([]BalanceChange, 0)

	query := []string{
		"SELECT COUNT(DISTINCT balance) - 1, coin, address",
		"FROM checkpoints",
		fmt.Sprintf("WHERE time > NOW() - INTERVAL '%d hours'", hours),
		"GROUP BY coin, address",
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
