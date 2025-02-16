package main

import (
	"log"
	"main/internal/handlers"
	"main/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// Инициализация роутера
	r := mux.NewRouter()

	// Регистрация обработчиков
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	// Статические файлы
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Запуск сервера
	utils.Logger.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}
}
