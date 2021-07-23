package cfg

import (
	"flag"

	"github.com/cs-bird/internals/psql"
)

type Cfg struct {
	PsqlCfg *psql.Config
	Migrate string
}

func Get() *Cfg {
	cfg := &Cfg{}

	flag.StringVar(&cfg.Migrate, "migrate", "up", "specify if we should be migrating DB 'up' or 'down'")
	cfg.PsqlCfg = psql.GetPGCfg()

	return cfg
}
