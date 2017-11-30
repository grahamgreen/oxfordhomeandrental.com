package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ohr.html")
	t.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	log.Fatal(http.ListenAndServe(":9000", r))
}
