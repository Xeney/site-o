package handlers

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Удаляем данные сессии
	session.Values["user"] = nil
	session.Values["isAuth"] = false
	session.Options.MaxAge = -1 // Удаляем cookie

	// Сохраняем сессию
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
