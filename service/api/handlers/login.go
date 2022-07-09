package handlers

import (
	"context"
	"encoding/json"
	"loadBalancedService/api/repository"
	"log"
	"net/http"
)

type loginHandler struct {
	repo *repository.API
}

func NewLoginHandler(repo *repository.API) *loginHandler {
	return &loginHandler{
		repo: repo,
	}
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.login(w, r)
		return
	}
	http.NotFound(w, r)

}

func (h *loginHandler) login(w http.ResponseWriter, r *http.Request) {
	var credentials repository.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	cookie, err := h.repo.LoginService.Login(ctx, credentials.Username, credentials.Password)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	log.Printf("User %s has logged in", credentials.Username)
	http.SetCookie(w, cookie)
	_, _ = w.Write([]byte("Logged in"))

}
