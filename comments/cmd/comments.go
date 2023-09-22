package main

import (
	"comments/pkg/api"
	db "comments/pkg/postgres"
	"fmt"
	"log"
	"net/http"
)

const db_url string = "postgres://postgres:password@localhost:5432/apigateway?sslmode=disable"

func main() {
	log.Print("server has started")
	//start the db
	pgdb, err := db.New(db_url)
	if err != nil {
		log.Printf("error starting the database %v", err)
	}
	//get the router of the API by passing the db
	router := api.StartAPI(pgdb)
	//get the port from the environment variable
	port := "8081"
	//pass the router and start listening with the server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
