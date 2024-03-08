package main

import (
	"fmt"
	"log"

	"fight-club/internal/database"
	"fight-club/internal/server"
	"fight-club/pkg/env"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Coudn't load .env file")
	}

	// Database
	pgUrl := env.GetEnv("POSTGRES_URL")
	db := database.Connect(pgUrl)
	defer db.Close()

	// Server HTTP 
	port := env.GetEnv("SERVER_PORT")
	addr := fmt.Sprintf(":%s", port)

	server := server.NewServer()
	log.Fatal(server.Listen(addr))
}