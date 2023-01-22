package ports

import "github.com/teleporttacos/proto/pb"

// DBPort is the port for the adapter
type DBPort interface {
	CloseDBConnection()
	GetMenuItem(pb.MenuItem)
	PlaceOrder(pb.Order)
}
