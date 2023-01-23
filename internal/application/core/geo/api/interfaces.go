package api

import "github.com/teleporttacos/proto/pb"

// Geo
type Geo interface {
	CheckForZeroCoordinates(*pb.OrderRequest) error
}
