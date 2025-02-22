package handlers

import (
	"log"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := session.Values

	if r.Method == http.MethodGet {
		if session.Values["isAuth"] != true {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		err = templates.ExecuteTemplate(w, "profile.html", data)
		if err != nil {
			log.Printf("Ошибка рендеринга шаблона: %v", err)
			http.Error(w, "Ошибка рендеринга страницы", http.StatusInternalServerError)
		}
	}
}
