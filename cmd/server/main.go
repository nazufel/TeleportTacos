package main

import (
	"log"

	// application

	// adapters
	rpc "github.com/teleporttacos/internal/adapters/framework/left/grpc"
	"github.com/teleporttacos/internal/adapters/framework/left/http"
	"github.com/teleporttacos/internal/adapters/framework/right/scylla"
	"github.com/teleporttacos/internal/application/core/geo"
	"github.com/teleporttacos/internal/application/core/geo/api"
)

func main() {

	dbAdapter, err := scylla.NewAdapter()
	if err != nil {
		log.Fatalf("failed to connect to DB - %v", err)
	}

	defer dbAdapter.CloseDBConnection()

	dbAdapter.SeedDatabase()

	core := geo.New()

	appAPI := api.NewApplication(dbAdapter, core)

	httpAdapter := http.NewAdapter()

	go httpAdapter.Run()

	gRPCAdapter := rpc.NewAdapter(appAPI)
	gRPCAdapter.Run()
	// db, err := scylla.NewAdapter()
	// if err != nil {
	// 	log.Fatalf("cannot connect to scylla: %v", err)
	// }

	// defer db.CloseDBConnection()

	// menuItem, err := db.GetMenuItem(pb.MenuItem{})
	// if err != nil {
	// 	log.Printf("error: %v", err)
	// }

	// log.Printf("menu item - name: %v, description: %v, price: $%v", menuItem.Name, menuItem.Description, menuItem.Price)

	// newOrder := pb.Order{
	// 	Count:        1,
	// 	MenuItem:     "The Taco",
	// 	Price:        5.00,
	// 	TeleportAlt:  123.87,
	// 	TeleportLat: 33.45,
	// 	TeleportLong: 27.98,
	// }

	// err = db.PlaceOrder(newOrder)
	// if err != nil {
	// 	log.Printf("unable to place order: %v", err)
	// }
}
