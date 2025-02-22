package main

import (
	"log"
	"main/internal/database"
	"main/internal/handlers"
	"net/http"
)

func main() {
	database.InitDB("./badgerdb")
	handlers.RegisterRoutes()
	log.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
