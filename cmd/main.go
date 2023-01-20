package main

import (
	"log"

	"github.com/TeleportTacos/internal/adapters/framework/right/scylla"
)

func main() {

	adapter, err := scylla.NewAdapter()
	if err != nil {
		log.Fatalf("cannot connect to scylla: %v", err)
	}

	log.Printf("scylla adapter: %v", adapter)
}
