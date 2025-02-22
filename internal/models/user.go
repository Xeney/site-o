package models

import (
	"time"
)

type User struct {
	Login        string    // Логин (уникальный)
	PasswordHash string    // Хэш пароля
	Email        string    // Электронная почта (уникальная)
	CountGames   int       // Количество сыгранных игр
	CountVictory int       // Количество побед
	CountDefeat  int       // Количество поражений
	Points       int       // Рейтинговые баллы
	Friends      []string  // Список идентификаторов друзей
	CreatedAt    time.Time // Дата регистрации
	LastActive   time.Time // Дата последней активности
	Role         string    // Роль пользователя (например, "user", "admin")
}
