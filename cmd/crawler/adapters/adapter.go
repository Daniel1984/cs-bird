package adapters

import "github.com/cs-bird/internals/models"

type Adapter interface {
	Get() (models.Checkpoint, error)
}
