package handlers

import (
	"fmt"
	"loadBalancedService/api/repository"
	"net/http"
)

type greetingHandler struct {
	port string // just for distinguishing between instances
	repo *repository.API
}

func NewGreetingHandler(port string, repo *repository.API) *greetingHandler {
	return &greetingHandler{port, repo}
}

func (h *greetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(fmt.Sprintf("Hello!, This instance is running on port: %s", h.port)))
}
