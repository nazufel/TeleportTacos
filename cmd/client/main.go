package main

import (
	"context"
	"log"
	"time"

	"github.com/teleporttacos/proto/pb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("unable to connect to server: %v", err)
	}

	log.Printf("dialed to server")

	defer conn.Close()

	c := pb.NewTacoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuItem, err := c.GetMenuItem(ctx, &pb.MenuItemRequest{})
	if err != nil {
		log.Printf("unable to get menu item from server: %v", err)
	}

	log.Printf("menu item: %v", menuItem)

	order := pb.OrderRequest{
		Count:        1,
		MenuItem:     menuItem.Name,
		Payment:      5.00,
		TeleportAlt:  54.23,
		TeleportLat:  99.99,
		TeleportLong: 100.1,
	}

	req, err := c.PlaceOrder(ctx, &order)
	if err != nil {
		log.Printf("unable to place order: %v", err)
	}

	log.Printf("order response: %v", req)
}
