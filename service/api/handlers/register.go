package handlers

import (
	"encoding/json"
	"fmt"
	"loadBalancedService/api/repository"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type registerHandler struct {
	port string // just for distinguishing between instances
	repo *repository.API
}

func NewRegisterHandler(port string, repo *repository.API) *registerHandler {
	return &registerHandler{
		port: port,
		repo: repo,
	}
}

func (h *registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.register(w, r)
	}

}

func (h *registerHandler) register(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		return
	}

	err = h.repo.RegisterService.Register(credentials.Username, credentials.Password)
	if err != nil {
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Username: %s has been created successfully", credentials.Username)))
}
