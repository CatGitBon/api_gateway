package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"

	"microservices/api_gateway/internal/handlers"
	authService "microservices/auth_service/pkg"
)

// Конфигурация для гRPC клиентов
var authClient authService.AuthServiceClient

// Инициализация gRPC клиентов
func initGRPCClients() {
	connAuth, err := grpc.Dial("auth_service:50051", grpc.WithInsecure()) // auth-service - это имя контейнера
	if err != nil {
		log.Fatalf("Failed to connect to Auth service: %v", err)
	}
	authClient = authService.NewAuthServiceClient(connAuth)

	handlers.SetAuthClient(authClient)
}

func main() {
	// Инициализируем gRPC-клиенты
	initGRPCClients()

	// Создаем роутер
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Здесь будут маршруты для обработки запросов от клиента
	r.Get("/auth", handlers.Authenticate)

	// Запуск HTTP-сервера
	log.Println("Starting API Gateway on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
