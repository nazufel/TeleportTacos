package ports

import "github.com/teleporttacos/proto/pb"

// DBPort is the port for the adapter
type DBPort interface {
	CloseDBConnection()
	GetMenuItem(*pb.MenuItemRequest) (pb.MenuItemResponse, error)
	PlaceOrder(*pb.OrderRequest) (pb.OrderResponse, error)
}
