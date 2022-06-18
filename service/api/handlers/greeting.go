package handlers

import (
	"fmt"
	"net/http"
)

type greetingHandler struct {
	port string // just for distinguishing between instances
}

func (h *greetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(fmt.Sprintf("Hello!, This instance is running on port: %s", h.port)))
}

func NewGreetingHandler(port string) *greetingHandler {
	return &greetingHandler{port}
}
