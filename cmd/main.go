package main

import (
	"log"

	"github.com/TeleportTacos/internal/adapters/framework/right/scylla"
)

func main() {

	db, err := scylla.NewAdapter()
	if err != nil {
		log.Fatalf("cannot connect to scylla: %v", err)
	}

	defer db.CloseDBConnection()

	db.SeedDatabase()

}
