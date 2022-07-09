package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"loadBalancedService/api/repository"
	"log"
	"net/http"
)

type registerHandler struct {
	repo *repository.API
}

func NewRegisterHandler(repo *repository.API) *registerHandler {
	return &registerHandler{
		repo: repo,
	}
}

func (h *registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.register(w, r)
		return
	}
	http.NotFound(w, r)

}

func (h *registerHandler) register(w http.ResponseWriter, r *http.Request) {
	var credentials repository.Credentials
	ctx := context.Background()

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		return
	}

	err = h.repo.RegisterService.Register(ctx, credentials.Username, credentials.Password)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Username: %s has been created successfully", credentials.Username)))
}
