// geo package holds core business logic concerning geographic rules

package geo

import (
	"errors"

	"github.com/teleporttacos/proto/pb"
)

// Geo struct implements the Geo interface
type Geo struct{}

// New returns a new Geo object
func New() *Geo {
	return &Geo{}
}

var ZeroCoordinatesError = errors.New("passed coordinates are: 0.0, 0.0, 0.0. Cannot teleport tacos to this location.")

// CheckForZeroCoordinates check if the passed coordinate passed are zero lat, long, altatude
// business rule doesn't not allow to teleport tacos to 0.0, 0.0, 0.0
func (g Geo) CheckForZeroCoordinates(o *pb.OrderRequest) error {
	if o.TeleportAlt == 0.0 && o.TeleportLat == 0.0 && o.TeleportLong == 0.0 {
		return ZeroCoordinatesError
	}

	return nil
}
