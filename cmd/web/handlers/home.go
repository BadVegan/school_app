package handlers

import (
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"net/http"
)

func Home(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from School App"))
	}
}
