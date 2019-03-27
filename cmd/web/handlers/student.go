package handlers

import (
	"encoding/json"
	"github.com/RyczkoDawid/school_app/cmd/web/helpers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"github.com/RyczkoDawid/school_app/pkg/models/mysql"
	"net/http"
	"strconv"
)

func GetStudent(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get(":id"))

		if err != nil || id < 1 {
			helpers.NotFound(app, w, err)
			return
		}

		s, err := app.Student.Get(id)
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
func GetStudents(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := app.Student.GetAllStudents()

		if err != nil {
			helpers.ServerError(app, w, err)
		}

		json.NewEncoder(w).Encode(s)
	}
}

func GetStudentsByClass(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get(":id"))

		if err != nil || id < 1 {
			helpers.NotFound(app, w, err)
			return
		}

		s, err := app.Student.GetAllByClass(id)
		if err != nil {
			helpers.ServerError(app, w, err)
		}

		json.NewEncoder(w).Encode(s)
	}
}

func CreateStudent(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = &mysql.Student{}

		err := json.NewDecoder(r.Body).Decode(s)
		if err != nil {
			helpers.ClientError(app, w, http.StatusBadRequest, err)
			return
		}

		id, err := app.Student.Insert(s)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		s, err = app.Student.Get(id)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		json.NewEncoder(w).Encode(s)
	}
}

func UpdateStudent(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = &mysql.Student{}

		err := json.NewDecoder(r.Body).Decode(s)
		if err != nil {
			helpers.ClientError(app, w, http.StatusBadRequest, err)
			return
		}

		err = app.Student.Update(s)
		if err != nil {
			helpers.ServerError(app, w, err)
			return
		}

		//s, err = app.Student.Get(id)
		//if err != nil {
		//	helpers.ServerError(app, w, err)
		//	return
		//}

		json.NewEncoder(w).Encode(s)
	}
}
