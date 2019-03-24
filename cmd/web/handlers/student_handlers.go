package handlers

import (
	"encoding/json"
	"github.com/RyczkoDawid/school_app/cmd/web/helpers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"github.com/RyczkoDawid/school_app/pkg/models"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"strconv"
)

func GetStudent(app *structs.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			helpers.NotFound(app, w, err)
			return
		}
		s, err := app.Student.Get(id)
		if err == models.ErrNoRecord {
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
		id, err := strconv.Atoi(r.URL.Query().Get("classId"))
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
		if r.Method != "POST" {
			w.Header().Set("Content-Type", "application/json")
			helpers.ClientError(app, w, http.StatusMethodNotAllowed, nil)
			return
		}
		s := &models.Student{
			Name:    "Dawid",
			Surname: "Ryczko",
			Phone:   "+48666666666",
			Email:   "davit@wp.pl",
			ClassID: null.NewInt(0, false),
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

		w.Header().Set("Allow", "POST")
		json.NewEncoder(w).Encode(s)
	}
}
