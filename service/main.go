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
	server := api.NewAPI(port)
	log.Printf("Server is listening at port: %s", port)
	log.Fatal(server.ListenAndServe())

}
