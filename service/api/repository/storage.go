package repository

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"loadBalancedService/ent"
	"log"
)

type API struct {
	db *ent.Client
}

func NewAPI() *API {
	client, err := ent.Open("mysql", "root:root@tcp(localhost:3306)/db?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer client.Close()

	// Auto Migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Init services

	// Build Repository API
	api := API{
		db: client,
	}
	return &api
}

//
//func (a *API) verifyLogin(handler *http.Handler) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		cookie, _ := r.Cookie("username")
//		if cookie.Value == "amityahav" {
//			handle
//		}
//	}
//
//}
