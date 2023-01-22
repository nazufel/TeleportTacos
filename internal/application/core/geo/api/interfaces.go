package api

// Geo
type Geo interface {
	CheckForZeroCoordinates(alt, lat, long float32) error
}
