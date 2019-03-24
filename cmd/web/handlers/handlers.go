package handlers

import (
	"github.com/RyczkoDawid/school_app/cmd/web/helpers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"net/http"
)

func Home(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			helpers.NotFound(app, w, nil)
			return
		}
		w.Write([]byte("Hello from Snippetbox"))
	}
}
