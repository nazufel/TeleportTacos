package ports

import "github.com/teleporttacos/proto/pb"

type APIPort interface {
	GetMenuItem(pb.MenuItem)
	PlaceOrder(pb.Order)
}
