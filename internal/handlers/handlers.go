package handlers

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func RegisterRoutes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", GetHome)
}

func init() {
	var err error
	templates, err = template.ParseFiles(
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/home.html",
	)
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблонов: %v", err)
	}
}
