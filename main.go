package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"jwt-server/handlers"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET enviroment variable is not set")
	}

	handlerInstance := handlers.NewHandler([]byte(jwtSecret))

	http.HandleFunc("/login", handlerInstance.LoginHandler)


	port := ":8080"
	fmt.Printf("Server has started on port: %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server has dropped: %v", err)
	}
}