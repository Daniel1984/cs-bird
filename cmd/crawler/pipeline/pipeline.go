package pipeline

import (
	"github.com/cs-bird/cmd/crawler/adapters"
	"github.com/cs-bird/internals/models"
)

type Report struct {
	Err error
	Res models.Checkpoint
}

type Pipeline struct {
	Adapters []adapters.Adapter
}

func New() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Add(a adapters.Adapter) {
	p.Adapters = append(p.Adapters, a)
}

func (p *Pipeline) Process() map[string]Report {
	resp := map[string]Report{}

	for _, a := range p.Adapters {
		report, err := a.Get()
		resp[report.Coin] = Report{
			Err: err,
			Res: report,
		}
	}

	return resp
}
