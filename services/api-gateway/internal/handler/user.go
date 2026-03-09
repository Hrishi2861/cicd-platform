package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	ServiceURL string
	Client     *http.Client
}

func NewUserHandler(serviceURL string) *UserHandler {
	return &UserHandler{
		ServiceURL: serviceURL,
		Client:     &http.Client{},
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Client.Get(h.ServiceURL + "/users")
	if err != nil {
		http.Error(w, "Failed to reach user service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	resp, err := h.Client.Get(h.ServiceURL + "/users/" + id)
	if err != nil {
		http.Error(w, "Failed to reach user service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Client.Post(h.ServiceURL+"/users", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to reach user service", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "healthy",
		"service": "api-gateway",
	})
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/users":
		if r.Method == http.MethodGet {
			h.GetUsers(w, r)
		} else if r.Method == http.MethodPost {
			h.CreateUser(w, r)
		}
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
