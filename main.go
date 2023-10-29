package main

import (
	"net/http"
	"github.com/gorilla/mux"
  "html/template"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, nil)
}
