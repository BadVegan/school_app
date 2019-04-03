package helpers

import (
	"fmt"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func ServerError(app *structs.Application, w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func ClientError(app *structs.Application, w http.ResponseWriter, status int, err error) {
	app.ErrorLog.Output(3, err.Error())
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func NotFound(app *structs.Application, w http.ResponseWriter, err error) {
	ClientError(app, w, http.StatusNotFound, err)
}
