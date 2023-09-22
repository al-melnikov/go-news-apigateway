package main

import (
	"censor/pkg/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("server has started")
	//get the router of the API by passing the db
	router := api.StartAPI()
	//get the port from the environment variable
	port := "8083"
	//pass the router and start listening with the server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
