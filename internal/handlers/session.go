package handlers

import (
	"encoding/gob"
	"main/internal/models"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

func init() {
	// Регистрируем тип User в gob
	gob.Register(models.User{})
}
