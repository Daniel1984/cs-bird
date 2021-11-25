package getbalancechange

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/boilerplate/pkg/application"
	"github.com/boilerplate/pkg/middleware"
	"github.com/julienschmidt/httprouter"
)

func getBalanceChangeEvents(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		qs := r.URL.Query()
		_, _ := strconv.Atoi(qs.Get("hours"))

		// cp := models.Checkpoint{}
		// cps, err := cp.

		// if err := models.GetByID(r.Context(), app); err != nil {
		// 	if errors.Is(err, sql.ErrNoRows) {
		// 		w.WriteHeader(http.StatusPreconditionFailed)
		// 		fmt.Fprintf(w, "user does not exist")
		// 		return
		// 	}

		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	fmt.Fprintf(w, "Oops")
		// 	return
		// }

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(nil)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		validateRequest,
	}

	return middleware.Chain(getBalanceChangeEvents(app), mdw...)
}
