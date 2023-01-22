package main

import (
	"log"

	"github.com/teleporttacos/internal/adapters/framework/right/scylla"
)

func main() {

	db, err := scylla.NewAdapter()
	if err != nil {
		log.Fatalf("cannot connect to scylla: %v", err)
	}

	defer db.CloseDBConnection()

	db.SeedDatabase()

	menuItem, err := db.GetMenuItem()
	if err != nil {
		log.Printf("error: %v", err)
	}

	log.Printf("menu item - name: %v, description: %v, price: $%v", menuItem.Name, menuItem.Description, menuItem.Price)

}
