package handlers

import (
	"log"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Printf("Ошибка рендеринга шаблона: %v", err)
		http.Error(w, "Ошибка рендеринга страницы", http.StatusInternalServerError)
	}
}
