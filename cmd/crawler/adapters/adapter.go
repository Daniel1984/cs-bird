package adapters

import "github.com/cs-bird/cmd/crawler/types"

type Adapter interface {
	Get() (types.Checkpoint, error)
}
