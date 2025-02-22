package handlers

import (
	"fmt"
	"log"
	"main/internal/database"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.Values["isAuth"] == true {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	data := session.Values

	if r.Method == http.MethodGet {
		err := templates.ExecuteTemplate(w, "register.html", data)
		if err != nil {
			log.Printf("Ошибка рендеринга шаблона: %v", err)
			http.Error(w, "Ошибка рендеринга страницы", http.StatusInternalServerError)
		}
	}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Ошибка при чтении данных", http.StatusBadRequest)
			return
		}
		// Получаем данные из формы
		login := r.FormValue("login")
		email := r.FormValue("email")
		password := r.FormValue("password")

		err = database.RegisterUser(login, password, email)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
