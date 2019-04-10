package handlers

import (
	"encoding/json"
	"github.com/RyczkoDawid/school_app/cmd/web/helpers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"github.com/RyczkoDawid/school_app/pkg/models/mysql"
	"net/http"
	"strconv"
)

func GetClass(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get(":id"))

		if err != nil || id < 1 {
			helpers.NotFound(app, w, err)
			return
		}

		c, err := app.Class.Get(id)

		if err == mysql.ErrNoRecord {
			helpers.NotFound(app, w, err)
			return

		} else if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}

func GetClasses(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := app.Class.GetClasses()

		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}
func CreateClass(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c = &mysql.Class{}

		err := json.NewDecoder(r.Body).Decode(c)
		if err != nil {
			helpers.ClientError(app, w, http.StatusBadRequest, err)
			return
		}

		id, err := app.Class.Insert(c)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		c, err = app.Class.Get(id)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}

func UpdateClass(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c = &mysql.Class{}

		err := json.NewDecoder(r.Body).Decode(c)
		if err != nil {
			helpers.ClientError(app, w, http.StatusBadRequest, err)
			return
		}

		err = app.Class.Update(c)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}
