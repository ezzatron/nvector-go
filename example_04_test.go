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
	// Position B is given with lat, long and height:
	bLat := 1.0
	bLon := 2.0
	bHeight := 3.0

	// Find the vector p_EB_E ("ECEF-vector")

	// SOLUTION:

	// Step1: Convert to n-vector:
	nvb := nvector.FromLatLon(nvector.Rad(bLat), nvector.Rad(bLon))

	// Step2: Find the ECEF-vector p_EB_E:
	pb := nvector.ToECEF(nvb, -bHeight)

	fmt.Printf("p_EB_E = [%.8f, %.8f, %.8f] m\n", pb.X, pb.Y, pb.Z)

	// Output:
	// p_EB_E = [6373290.27721828, 222560.20067474, 110568.82718179] m
}
