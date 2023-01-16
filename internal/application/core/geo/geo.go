// geo package holds core business logic concerning geographic rules

package geo

import "errors"

// Geo struct implements the Geo interface
type Geo struct{}

// New returns a new Geo object
func New() *Geo {
	return &Geo{}
}

var ZeroCoordinatesError = errors.New("passed coordinates are: 0.0, 0.0, 0.0. Cannot teleport tacos to this location.")

// CheckForZeroCoordinates check if the passed coordinate passed are zero lat, long, altatude
// business rule doesn't not allow to teleport tacos to 0.0, 0.0, 0.0
func (g Geo) CheckForZeroCoordinates(lat, long, alt float32) error {
	if lat == 0.0 && long == 0.0 && alt == 0.0 {
		return ZeroCoordinatesError
	}

	return nil
}
