package main

import (
	. "github.com/RyczkoDawid/school_app/cmd/web/handlers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"github.com/bmizerany/pat"
	"net/http"
)

func routes(app *structs.Application) *http.ServeMux {
	mux := pat.New()
	mux.Get("/", Home(app))
	mux.Get("/class", GetClass(app))
	mux.Get("/classes", GetClasses(app))
	mux.Post("/class/create", CreateClass(app))
	mux.Get("/student/:id", GetStudent(app))
	mux.Put("/student/:id", UpdateStudent(app))
	mux.Get("/students/class/:id", GetStudentsByClass(app))
	mux.Get("/students/all", GetStudents(app))
	mux.Post("/student/create", CreateStudent(app))

	sm := http.NewServeMux()
	sm.Handle("/", mux)
	return sm
}
