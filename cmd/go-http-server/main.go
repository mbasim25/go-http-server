package main

import (
	"log"
	"net/http"

	"github.com/mbasim25/go-http-server/postgres"
	"github.com/mbasim25/go-http-server/web"
)

func main() {
	store, err := postgres.NewStore(
		"postgres://postgres:password@localhost/postgres?sslmode=disable", // TODO: add env variables
	)
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
