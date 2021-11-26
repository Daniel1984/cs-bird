package middleware

import (
	"github.com/julienschmidt/httprouter"
)

// Middleware type
type Middleware func(httprouter.Handle) httprouter.Handle

// Chain - chains all middleware functions
func Chain(f httprouter.Handle, m ...Middleware) httprouter.Handle {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise run recursively over nested handlers
	return m[0](Chain(f, m[1:]...))
}
