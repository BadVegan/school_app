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
	mux.Get("/students", GetStudents(app))
	mux.Post("/student/create", CreateStudent(app))

	mux.Get("/summaryLesson/:id", GetSummaryLesson(app))

	mux.Get("/teacher/:id", GetTeacher(app))
	mux.Put("/teacher/:id", UpdateTeacher(app))
	mux.Get("/teachers", GetTeachers(app))
	mux.Post("/teacher/create", CreateTeacher(app))

	sm := http.NewServeMux()
	sm.Handle("/", mux)
	return sm
}
