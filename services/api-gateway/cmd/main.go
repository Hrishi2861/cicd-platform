package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cicd-platform/services/api-gateway/internal/handler"
	"github.com/cicd-platform/services/api-gateway/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter()

	router.Use(middleware.MetricsMiddleware)
	router.Use(middleware.LoggingMiddleware)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	userServiceURL := getEnv("USER_SERVICE_URL", "http://localhost:8081")
	orderServiceURL := getEnv("ORDER_SERVICE_URL", "http://localhost:8082")

	userHandler := handler.NewUserHandler(userServiceURL)
	orderHandler := handler.NewOrderHandler(orderServiceURL)

	apiRouter.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	apiRouter.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	apiRouter.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	apiRouter.HandleFunc("/orders", orderHandler.GetOrders).Methods("GET")
	apiRouter.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
	apiRouter.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")

	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())

	port := getEnv("PORT", "8080")
	log.Printf("API Gateway starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
