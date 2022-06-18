package api

import (
	"fmt"
	"loadBalancedService/api/handlers"
	"net/http"
)

func NewAPI(port string) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/api", handlers.NewGreetingHandler(port))

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return &server
}
