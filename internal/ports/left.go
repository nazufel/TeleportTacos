package ports

import "github.com/teleporttacos/proto/pb"

type APIPort interface {
	GetMenuItem(*pb.MenuItemRequest) (pb.MenuItemResponse, error)
	PlaceOrder(*pb.OrderRequest) (pb.OrderResponse, error)
}
