package database

import (
	"encoding/json"
	"errors"
	"main/internal/models"
	"regexp"
	"time"

	"github.com/dgraph-io/badger/v3"
	"golang.org/x/crypto/bcrypt"
)

// SaveUserByLogin сохраняет данные пользователя по логину.
func SaveUserByLogin(login string, userData []byte) error {
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte("user:login:"+login), userData)
	})
	return err
}

func checkPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUserByLogin возвращает данные пользователя по логину.
func GetUserByLogin(login string) ([]byte, error) {
	var userData []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("user:login:" + login))
		if err != nil {
			return err
		}
		userData, err = item.ValueCopy(nil)
		return err
	})
	return userData, err
}

// DeleteUserByLogin удаляет данные пользователя по логину.
func DeleteUserByLogin(login string) error {
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte("user:login:" + login))
	})
	return err
}

// IsLoginUnique проверяет, уникален ли логин.
func IsLoginUnique(login string) (bool, error) {
	_, err := GetUserByLogin(login)
	if err == badger.ErrKeyNotFound {
		return true, nil // Логин уникален
	}
	if err != nil {
		return false, err // Произошла ошибка
	}
	return false, nil // Логин уже существует
}

// ValidateUserData проверяет данные пользователя на валидность.
func ValidateUserData(login, password, email string) error {
	// Проверка логина
	if len(login) < 3 || len(login) > 20 {
		return errors.New("логин должен быть от 3 до 20 символов")
	}

	// Проверка пароля
	if len(password) < 8 {
		return errors.New("пароль должен быть не короче 8 символов")
	}

	// Проверка email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("неверный формат email")
	}

	return nil
}

// RegisterUser регистрирует нового пользователя.
func RegisterUser(login, password, email string) error {
	// Проверка данных на валидность
	err := ValidateUserData(login, password, email)
	if err != nil {
		return err
	}

	// Проверка уникальности логина
	isUnique, err := IsLoginUnique(login)
	if err != nil {
		return err
	}
	if !isUnique {
		return errors.New("логин уже занят")
	}

	// Хэширование пароля
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	// Создание пользователя
	user := models.User{
		Login:        login,
		PasswordHash: hashedPassword,
		Email:        email,
		CountGames:   0,          // Начальное значение
		CountVictory: 0,          // Начальное значение
		CountDefeat:  0,          // Начальное значение
		Points:       0,          // Начальное значение
		Friends:      []string{}, // Пустой список друзей
		CreatedAt:    time.Now(), // Текущее время
		LastActive:   time.Now(), // Текущее время
		Role:         "user",     // Роль по умолчанию
	}

	// Сериализация пользователя в JSON
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Сохранение пользователя в BadgerDB
	err = SaveUserByLogin(login, userData)
	if err != nil {
		return err
	}

	return nil
}

func AuthUser(login, password string) (*models.User, error) {
	var user models.User
	// Проверка уникальности логина
	isUnique, err := IsLoginUnique(login)
	if err != nil {
		return &models.User{}, err
	}
	if isUnique {
		return &models.User{}, errors.New("Нет такого пользователя")
	}
	userData, err := GetUserByLogin(login)
	json.Unmarshal(userData, &user)
	if pass := checkPassword(user.PasswordHash, password); !pass {
		return &models.User{}, errors.New("Неверный пароль")
	}
	return &user, nil
}

// HashPassword хэширует пароль.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
