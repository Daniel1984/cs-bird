package getbalancechange

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func validateRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		qs := r.URL.Query()
		hoursStr := qs.Get("hours")

		if len(hoursStr) == 0 {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "?hours=x query string is required")
			return
		}

		h, err := strconv.Atoi(hoursStr)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "malformed `hours` query string, only accepts integers")
			return
		}

		if h < 1 || h > 48 {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "`hours` value mus be in range [1..24]")
			return
		}

		next(w, r, p)
	}
}
