package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/teleporttacos/proto/pb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("api:9999", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to connect to server: %v", err)
	}

	log.Printf("dialed to server")

	defer conn.Close()

	c := pb.NewTacoServiceClient(conn)

	// generate load on the system
	for {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		menuItem, err := c.GetMenuItem(ctx, &pb.MenuItemRequest{})
		if err != nil {
			log.Printf("unable to get menu item from server: %v", err)
		}

		log.Printf("available menu item: %v", menuItem.Name)

		coordinates := randCoordinates()

		order := pb.OrderRequest{
			Count:        1,
			MenuItem:     menuItem.Name,
			Payment:      5.00,
			TeleportAlt:  coordinates[0],
			TeleportLat:  coordinates[1],
			TeleportLong: coordinates[2],
		}

		_, err = c.PlaceOrder(ctx, &order)
		if err != nil {
			log.Printf("unable to place order: %v", err)
		}

		log.Println("order placed")

		// time.Sleep(100 * time.Millisecond)
	}

}

func randCoordinates() []float32 {
	var min float32
	var max float32

	min = -180.00
	max = 180.00

	rand.Seed(time.Now().UnixNano())
	c := make([]float32, 3)
	for i := range c {
		c[i] = min + rand.Float32()*(max-min)
	}

	return c
}
