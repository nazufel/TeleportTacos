package main

import rpc "github.com/teleporttacos/internal/adapters/framework/left/grpc"

func main() {

	// db, err := scylla.NewAdapter()
	// if err != nil {
	// 	log.Fatalf("cannot connect to scylla: %v", err)
	// }

	// defer db.CloseDBConnection()

	// db.SeedDatabase()

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
	// 	TeleportLang: 33.45,
	// 	TeleportLong: 27.98,
	// }

	// err = db.PlaceOrder(newOrder)
	// if err != nil {
	// 	log.Printf("unable to place order: %v", err)
	// }

	gRPCAdapter := rpc.NewAdapter()
	gRPCAdapter.Run()
}
