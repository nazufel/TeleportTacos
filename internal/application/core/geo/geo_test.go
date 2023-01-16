package geo

import "testing"

// test passing zeros
func TestCheckForZeroCoordinates(t *testing.T) {

	geo := New()

	// case pass all zeros to the core
	err := geo.CheckForZeroCoordinates(0.0, 0.0, 0.0)
	if err != ZeroCoordinatesError {
		t.Fatalf("expected to error since passed: 0.0, 0.0, 0.0")
	}

	// case don't pass all zeros to the core
	err = geo.CheckForZeroCoordinates(1.0, 1.0, 1.0)
	if err != nil {
		t.Fatalf("did not expect to fail since passing non-zeros")
	}

}
