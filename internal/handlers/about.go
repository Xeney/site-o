package handlers

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/about.html"))
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
