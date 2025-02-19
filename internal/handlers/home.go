package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/home.html"))

	tmpl.ExecuteTemplate(w, "home.html", nil)
}
