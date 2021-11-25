package application

import (
	"flag"
	"os"

	"github.com/cs-bird/internals/psql"
)

// Application holds commonly used app wide data, for ease of DI
type Application struct {
	DB      *psql.DB
	ApiPort string
}

// Get captures env vars, establishes DB connection and returns reference to both
func Get() (*Application, error) {
	app := &Application{}

	flag.StringVar(&app.ApiPort, "apiPort", os.Getenv("API_PORT"), "API Port")

	db, err := psql.New()
	if err != nil {
		return nil, err
	}

	app.DB = db
	return app, nil
}
