package main

import (
	"log"

	"github.com/Thalisonh/crud-golang/server"
	"github.com/joho/godotenv"
)

func main() {
	errDotEnv := godotenv.Load()

	if errDotEnv != nil {
		log.Fatal("Error loading .env files")
	}

	server := server.NewServer()

	server.Run()

}
