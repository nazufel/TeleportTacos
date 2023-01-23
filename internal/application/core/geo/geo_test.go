package geo

import (
	"testing"

	"github.com/teleporttacos/proto/pb"
)

// test passing zeros
func TestCheckForZeroCoordinates(t *testing.T) {

	geo := New()

	or1 := &pb.OrderRequest{
		TeleportAlt:  0.0,
		TeleportLat:  0.0,
		TeleportLong: 0.0,
	}

	// case pass all zeros to the core
	err := geo.CheckForZeroCoordinates(or1)
	if err != ZeroCoordinatesError {
		t.Fatalf("expected to error since passed: 0.0, 0.0, 0.0")
	}

	or2 := &pb.OrderRequest{
		TeleportAlt:  1.0,
		TeleportLat:  1.0,
		TeleportLong: 1.0,
	}

	// case don't pass all zeros to the core
	err = geo.CheckForZeroCoordinates(or2)
	if err != nil {
		t.Fatalf("did not expect to fail since passing non-zeros")
	}

}
