package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Загрузка шаблона
	tmpl := template.Must(template.ParseFiles("web/templates/home.html"))

	// Данные для шаблона
	data := struct {
		Title string
	}{
		Title: "Добро пожаловать!",
	}

	// Рендеринг шаблона
	tmpl.ExecuteTemplate(w, "home.html", data)
}
