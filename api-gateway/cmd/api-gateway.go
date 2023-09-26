package main

import (
	"api-gateway/pkg/api"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	//get the port from the environment variable
	portString := os.Getenv("GATEWAY_PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	log.Print("server has started")
	//get the router of the API by passing the db
	router := api.StartAPI()

	//pass the router and start listening with the server
	err := http.ListenAndServe(fmt.Sprintf(":%s", portString), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
