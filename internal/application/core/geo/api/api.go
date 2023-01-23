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

// GetMenuItem gets a menu item
func (a Application) GetMenuItem(m *pb.MenuItemRequest) (*pb.MenuItemResponse, error) {

	// TODO: handle error and respond with grpc error
	menuItem, _ := a.db.GetMenuItem(m)

	res := pb.MenuItemResponse{
		Id:          menuItem.Id,
		Name:        menuItem.Name,
		Description: menuItem.Description,
		Price:       menuItem.Price,
	}

	return &res, nil
}

// PlaceOrder gets a menu item
func (a Application) PlaceOrder(m *pb.OrderRequest) (*pb.OrderResponse, error) {

	var res pb.OrderResponse

	// validate order has valid teleportation coordinates
	err := a.geo.CheckForZeroCoordinates(m)
	if err != nil {
		log.Printf("requested order does not have valid coordinates. received alt: %v, lat: %v, long: %v", m.TeleportAlt, m.TeleportLat, m.TeleportLong)

		return &res, err
	}

	// valid coordinates can place an order in the db
	order, err := a.db.PlaceOrder(m)
	if err != nil {
		log.Printf("error getting record from db: %v", err)
		return &res, err
	}

	return order, err
}
