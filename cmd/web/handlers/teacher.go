package handlers

import (
	"encoding/json"
	"github.com/RyczkoDawid/school_app/cmd/web/helpers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"github.com/RyczkoDawid/school_app/pkg/models/mysql"
	"net/http"
	"strconv"
)

func CreateTeacher(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = &mysql.Teacher{}

		err := json.NewDecoder(r.Body).Decode(s)
		if err != nil {
			helpers.ClientError(app, w, http.StatusBadRequest, err)
			return
		}

		id, err := app.Teacher.Insert(s)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		s, err = app.Teacher.Get(id)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(s)
	}
}

func GetTeacher(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get(":id"))

		if err != nil || id < 1 {
			helpers.NotFound(app, w, err)
			return
		}

		s, err := app.Teacher.Get(id)

		if err == mysql.ErrNoRecord {
			helpers.NotFound(app, w, err)
			return

		} else if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(s)
	}
}

func GetTeachers(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := app.Teacher.GetTeachers()

		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(s)
	}
}

func UpdateTeacher(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = &mysql.Teacher{}

		err := json.NewDecoder(r.Body).Decode(s)
		if err != nil {
			helpers.ClientError(app, w, http.StatusBadRequest, err)
			return
		}

		err = app.Teacher.Update(s)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(s)
	}
}
