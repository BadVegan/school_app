package main

import (
	. "github.com/RyczkoDawid/school_app/cmd/web/handlers"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"net/http"
)

func routes(app *structs.Application) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", Home(app))
	mux.Handle("/class", GetClass(app))
	mux.Handle("/classes", GetClasses(app))
	mux.Handle("/class/create", CreateClass(app))
	mux.Handle("/student", GetStudent(app))
	mux.Handle("/students", GetStudentsByClass(app))
	mux.Handle("/students/all", GetStudents(app))
	mux.Handle("/student/create", CreateStudent(app))
	return mux
}
