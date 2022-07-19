package api

import (
	"fmt"
	"loadBalancedService/api/handlers"
	"loadBalancedService/api/repository"
	"net/http"
)

func NewAPI(port string) *http.Server {
	repoAPI := repository.NewAPI()

	// Handlers
	greetingHandler := handlers.NewGreetingHandler(port, repoAPI)
	authHandler := handlers.NewAuthHandler()

	// Routes
	mux := http.NewServeMux()
	mux.Handle("/", authHandler.ValidateJWT(greetingHandler))
	mux.Handle("/get-token", authHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return &server
}
