package handlers

import (
	"log"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := session.Values

	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	err = templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		log.Printf("Ошибка рендеринга шаблона: %v", err)
		http.Error(w, "Ошибка рендеринга страницы", http.StatusInternalServerError)
	}
}
