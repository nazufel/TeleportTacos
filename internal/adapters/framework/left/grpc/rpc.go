package rpc

import (
	"context"
	"log"

	"github.com/teleporttacos/proto/pb"
)

// GetMenuItem wires up the api method
func (a Adapter) GetMenuItem(ctx context.Context, req *pb.MenuItemRequest) (*pb.MenuItemResponse, error) {

	res, err := a.api.GetMenuItem(req)
	if err != nil {
		log.Printf("error talking to the api layer for get menu item: %v", err)
		return &res, nil
	}

	return &res, nil
}

// PlaceOrder wires up the api method
func (a Adapter) PlaceOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {

	res, err := a.api.PlaceOrder(req)
	if err != nil {
		log.Printf("error talking to the api layer to place an order: %v", err)
		return &res, nil
	}

	return &res, nil
}
