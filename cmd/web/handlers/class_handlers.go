package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetClass(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific class with ID %d...", id)
}
func GetClasses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from getClasses"))
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"Alex"}`))
}
