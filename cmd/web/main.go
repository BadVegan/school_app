package main

import (
	"flag"
	"github.com/RyczkoDawid/school_app/cmd/web/handlers"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "localhost:5000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/class", handlers.GetClass)
	mux.HandleFunc("/classes", handlers.GetClasses)
	mux.HandleFunc("/class/create", handlers.CreateClass)
	mux.HandleFunc("/student", handlers.GetStudent)
	mux.HandleFunc("/students", handlers.GetStudents)
	mux.HandleFunc("/student/create", handlers.CreateStudent)

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
