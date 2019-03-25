package handlers

import (
	"fmt"
	"github.com/RyczkoDawid/school_app/cmd/web/helpers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"net/http"
	"strconv"
)

func GetClass(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			helpers.NotFound(app, w, err)
			return
		}
		fmt.Fprintf(w, "Display a specific class with ID %d...", id)
	}
}

func GetClasses(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from getClasses"))
	}
}
func CreateClass(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			helpers.ClientError(app, w, http.StatusMethodNotAllowed, nil)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name":"Alex"}`))
	}
}
