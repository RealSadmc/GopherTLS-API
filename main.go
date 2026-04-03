package main

import (
	"github.com/RealSadmc/GopherTLS-API/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	serverHost := os.Getenv("SERVER_HOST")

	if serverHost == "" {
		log.Fatal("SERVER_HOST not set, please check your env file!")
	}

	serverPort := os.Getenv("SERVER_PORT")

	if serverPort == "" {
		log.Fatal("SERVER_PORT not set, please check your env file!")
	}

	if err := server.StartServer(serverHost, serverPort); err != nil {
		log.Fatal(err)
	}
}
