package main

import (
	"github.com/Thalisonh/crud-golang/server"
	"github.com/Thalisonh/crud-golang/server/database"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}