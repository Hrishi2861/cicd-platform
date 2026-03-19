package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cicd-platform/services/order-service/internal/handler"
	"github.com/cicd-platform/services/order-service/internal/middleware"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "orders")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var db *sql.DB
	var err error

	for i := 0; i < 30; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			break
		}
		log.Printf("Waiting for database... (%d/30)", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after 30 retries: %v", err)
	}
	defer db.Close()

	log.Println("Connected to database successfully")

	orderHandler := handler.NewOrderHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			orderHandler.GetOrders(w, r)
		case http.MethodPost:
			orderHandler.CreateOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/orders/{id}", orderHandler.HandleOrderByID)
	mux.HandleFunc("/health", handler.HealthCheck)
	mux.Handle("/metrics", promhttp.Handler())

	loggedMux := middleware.LoggingMiddleware(mux)
	metricsMux := middleware.MetricsMiddleware(loggedMux)

	port := getEnv("PORT", "8082")
	log.Printf("Order Service starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, metricsMux))
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
