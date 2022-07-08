package main

import (
	"loadBalancedService/api"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Port number was not provided")
	}
	port := os.Args[1]
	logger := log.New(os.Stdout, "api", log.LstdFlags)
	server := api.NewAPI(port)
	logger.Printf("Server is listening at port: %s", port)
	logger.Fatal(server.ListenAndServe())

}
