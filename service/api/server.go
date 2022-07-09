package api

import (
	"fmt"
	"loadBalancedService/api/handlers"
	"loadBalancedService/api/repository"
	"net/http"
)

func NewAPI(port string) *http.Server {
	repoAPI := repository.NewAPI()

	// Routes
	mux := http.NewServeMux()
	mux.Handle("/", handlers.NewGreetingHandler(port, repoAPI))
	mux.Handle("/register", handlers.NewRegisterHandler(repoAPI))
	mux.Handle("/login", handlers.NewLoginHandler(repoAPI))

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return &server
}
