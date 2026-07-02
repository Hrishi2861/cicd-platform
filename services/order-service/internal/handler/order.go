package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OrderHandler struct {
	DB *sql.DB
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{DB: db}
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, user_id, product_name, quantity, total_price, status, created_at, updated_at FROM orders ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.ProductName, &o.Quantity, &o.TotalPrice, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		orders = append(orders, o)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) HandleOrderByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/orders/")
	if id == "" {
		http.Error(w, "Order ID required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetOrder(w, r, id)
	case http.MethodPut:
		h.UpdateOrder(w, r, id)
	case http.MethodDelete:
		h.DeleteOrder(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request, id string) {
	var o Order
	err := h.DB.QueryRow("SELECT id, user_id, product_name, quantity, total_price, status, created_at, updated_at FROM orders WHERE id = $1", id).
		Scan(&o.ID, &o.UserID, &o.ProductName, &o.Quantity, &o.TotalPrice, &o.Status, &o.CreatedAt, &o.UpdatedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(o)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID      string  `json:"user_id"`
		ProductName string  `json:"product_name"`
		Quantity    int     `json:"quantity"`
		TotalPrice  float64 `json:"total_price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	now := time.Now()

	_, err := h.DB.Exec(
		"INSERT INTO orders (id, user_id, product_name, quantity, total_price, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		id, input.UserID, input.ProductName, input.Quantity, input.TotalPrice, "pending", now, now,
	)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	order := Order{
		ID:          id,
		UserID:      input.UserID,
		ProductName: input.ProductName,
		Quantity:    input.Quantity,
		TotalPrice:  input.TotalPrice,
		Status:      "pending",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request, id string) {
	var input struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	now := time.Now()
	_, err := h.DB.Exec(
		"UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3",
		input.Status, now, id,
	)
	if err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	var o Order
	err = h.DB.QueryRow("SELECT id, user_id, product_name, quantity, total_price, status, created_at, updated_at FROM orders WHERE id = $1", id).
		Scan(&o.ID, &o.UserID, &o.ProductName, &o.Quantity, &o.TotalPrice, &o.Status, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(o)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request, id string) {
	_, err := h.DB.Exec("DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":  "healthy",
		"service": "order-service",
	})
}
