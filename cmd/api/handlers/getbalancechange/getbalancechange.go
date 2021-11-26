package getbalancechange

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cs-bird/cmd/api/application"
	"github.com/cs-bird/internals/middleware"
	"github.com/julienschmidt/httprouter"
)

func getBalanceChangeEvents(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		defer r.Body.Close()

		qs := r.URL.Query()
		hours, _ := strconv.Atoi(qs.Get("hours"))

		events, err := app.Repo.BalanceChange.FetchForHours(r.Context(), hours)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "failed getting balance change events: %s", err)
			return
		}

		res, err := json.Marshal(events)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "failed parsing json: %s", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(res); err != nil {
			log.Printf("failed responding to api call: %s\n", err)
		}
	}
}

func Do(app *application.Application) httprouter.Handle {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		validateRequest,
	}

	return middleware.Chain(getBalanceChangeEvents(app), mdw...)
}
