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
	// DB's URL should be `db:3306` when deploying everything via Docker.
	client, err := ent.Open("mysql", "root:root@tcp(db:3306)/db?parseTime=True")
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
