package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific student with ID %d...", id)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from getStudents"))
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Allow", "POST")
	w.Write([]byte(`{"name":"Alex"}`))
}
