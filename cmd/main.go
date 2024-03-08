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
		log.Println("Dotenv file is not loading environment variables")
	}

	// Database
	user := env.GetEnv("DB_USER")
	password := env.GetEnv("DB_PASSWORD")
	host := env.GetEnv("DB_HOST")
	port := env.GetEnv("DB_PORT")
	name := env.GetEnv("DB_NAME")
	pool := env.GetEnv("DB_POOL")
	db := database.Connect(user, password, host, port, name, pool)
	defer db.Close()

	// Server HTTP
	addr := env.GetEnv("HTTP_PORT")
	server := server.NewServer()
	log.Fatal(server.Listen(fmt.Sprintf(":%s", addr)))
}
