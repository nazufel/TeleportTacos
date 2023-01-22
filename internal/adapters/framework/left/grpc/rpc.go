package rpc

import (
	"context"

	"github.com/teleporttacos/proto/pb"
)

func (a Adapter) GetMenuItem(ctx context.Context, req *pb.MenuItemRequest) (*pb.MenuItemResponse, error) {

	var response pb.MenuItemResponse
	return &response, nil
}

func (a Adapter) PlaceOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {

	var response pb.OrderResponse
	return &response, nil
}
