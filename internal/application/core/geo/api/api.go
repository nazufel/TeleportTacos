package api

import (
	"log"

	"github.com/teleporttacos/internal/ports"
	"github.com/teleporttacos/proto/pb"
)

type Application struct {
	db  ports.DBPort
	geo Geo
}

func NewApplication(db ports.DBPort, geo Geo) *Application {
	return &Application{db: db, geo: geo}
}

// CheckForZeroCoordinates checks the business rules if we can teleport tacos to these coordinates
func (a Application) CheckForZeroCoordinates(alt, lat, long float32) error {

	err := a.geo.CheckForZeroCoordinates(alt, lat, long)
	if err != nil {
		log.Printf("unable to teleport tacos to alt: 0, lat: 0, long: 0")
	}

	return nil
}

// GetMenuItem gets a menu item
func (a Application) GetMenuItem(m *pb.MenuItemRequest) (*pb.MenuItemResponse, error) {

	var err error
	var res pb.MenuItemResponse

	return &res, err
}

// PlaceOrder gets a menu item
func (a Application) PlaceOrder(m *pb.OrderRequest) (*pb.OrderResponse, error) {

	var err error
	var res pb.OrderResponse

	return &res, err
}
