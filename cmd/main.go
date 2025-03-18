package main

import (
	"log"
	"github.com/dhichana/grad-connect/internal/database"
	"github.com/dhichana/grad-connect/internal/server"
)


func main() {
	db, err := database.InitDB("../database.db")
	if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
	if err := server.Start(db); err != nil{
		log.Fatalf("Failed to start server: %v", err)
	}
}
