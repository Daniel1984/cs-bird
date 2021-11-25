package router

import (
	"github.com/cs-bird/cmd/api/application"
	"github.com/cs-bird/cmd/api/handlers/getbalancechange"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/balance-change-events", getbalancechange.Do(app))
	return mux
}
