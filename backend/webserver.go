package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ExtraHandler struct {
	filename string
}

func (e *ExtraHandler) DogHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(ConvertTextToJson(e.filename))
	log.Println("Call to DogHandler")
	if err != nil {
		log.Fatalf("Error encoding the Dogs: %v", err)
	}
}
