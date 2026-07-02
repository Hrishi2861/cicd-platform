package handler

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	ServiceURL string
	Client     *http.Client
}

func NewOrderHandler(serviceURL string) *OrderHandler {
	return &OrderHandler{
		ServiceURL: serviceURL,
		Client:     &http.Client{},
	}
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Client.Get(h.ServiceURL + "/orders")
	if err != nil {
		http.Error(w, "Failed to reach order service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	resp, err := h.Client.Get(h.ServiceURL + "/orders/" + id)
	if err != nil {
		http.Error(w, "Failed to reach order service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Client.Post(h.ServiceURL+"/orders", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to reach order service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}
