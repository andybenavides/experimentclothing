package handler

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("main.html"))

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
