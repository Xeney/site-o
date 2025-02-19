package main

import (
	"fmt"
	"main/internal/handlers"
	"net/http"
)

func main() {
	handlers.RegisterRoutes()
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
