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
	http.HandleFunc("/login", Auth)
	http.HandleFunc("/reg", Register)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/profile", Profile)
}

func init() {
	var err error
	templates, err = template.ParseFiles(
		"web/templates/header.html",
		"web/templates/footer.html",
		"web/templates/home.html",
		"web/templates/register.html",
		"web/templates/auth.html",
		"web/templates/profile.html",
	)
	if err != nil {
		log.Fatalf("Ошибка загрузки шаблонов: %v", err)
	}
}
