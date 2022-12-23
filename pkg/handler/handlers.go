package handler

import (
	"html/template"
	"net/http"
)

func Create_Page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("webpage.html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
