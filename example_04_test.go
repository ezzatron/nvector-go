package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 4: Geodetic latitude to ECEF-vector
//
// Given geodetic latitude, longitude and height. Find the ECEF-vector (using
// WGS-84 ellipsoid).
//
// See: https://www.ffi.no/en/research/n-vector/#example_4
func Example_n04LatLonToECEF() {
	// PROBLEM:

	// Geodetic latitude, longitude and height are given for position B:
	bLat, bLon, bHeight := 1.0, 2.0, 3.0

	// Find the ECEF-vector for this position.

	// SOLUTION:

	// Step 1: First, the given latitude and longitude are converted to n-vector:
	b := nvector.Position{
		Vector: nvector.FromGeodeticCoordinates(
			nvector.GeodeticCoordinates{
				Latitude:  nvector.Radians(bLat),
				Longitude: nvector.Radians(bLon),
			},
			nvector.ZAxisNorth,
		),
		Depth: -bHeight,
	}

	// Step 2: Convert to an ECEF-vector:
	pb := nvector.ToECEF(b, nvector.WGS84, nvector.ZAxisNorth)

	fmt.Printf("p_EB_E = [%.8f, %.8f, %.8f] m\n", pb.X, pb.Y, pb.Z)

	// Output:
	// p_EB_E = [6373290.27721828, 222560.20067474, 110568.82718179] m
}
