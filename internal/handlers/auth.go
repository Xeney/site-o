package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/database"
	"net/http"
	"time"
)

func Auth(w http.ResponseWriter, r *http.Request) {
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
		err := templates.ExecuteTemplate(w, "auth.html", data)
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
		password := r.FormValue("password")

		user, err := database.AuthUser(login, password)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		userData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		user.LastActive = time.Now()
		database.SaveUserByLogin(login, userData)

		session.Values["isAuth"] = true
		session.Values["user"] = user

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/profile", http.StatusFound)
	}
}
