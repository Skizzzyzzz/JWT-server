package main

import (
	"fmt"
	"log"
	"net/http"
	"jwt-server/handlers"
)

//var jwtKey = []byte("very_secret_key")

const jwtSecret = "super-secret-key"

func main() {
	handlerInstance := handlers.NewHandler([]byte(jwtSecret))

	http.HandleFunc("/login", handlerInstance.LoginHandler)


	port := ":8080"
	fmt.Printf("Server has started on port: %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server has dropped: %v", err)
	}
}