package main

import (
	"log"
	"net/http"
	"os"
)

var (
	port = os.Getenv("BACKEND_PORT")
)

func main() {
	ListenAndServerPort := ":" + port
	eh := ExtraHandler{
		filename: "imaginaryDatabase.txt",
	}
	http.HandleFunc("/api/dogs", eh.DogHandler)
	log.Printf("Starting up the webserver at http://localhost:%s/api/dogs", port)
	err := http.ListenAndServe(ListenAndServerPort, nil)
	if err != nil {
		log.Fatalf("Error starting the application: %v", err)
	}
}
