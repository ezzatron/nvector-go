package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 3: ECEF-vector to geodetic latitude
//
// Given an ECEF-vector of a position. Find geodetic latitude, longitude and
// height (using WGS-84 ellipsoid).
//
// See: https://www.ffi.no/en/research/n-vector/#example_3
func Example_n03ECEFToLatLon() {
	// PROBLEM:

	// Position B is given as an “ECEF-vector” pb (i.e. a vector from E, the
	// center of the Earth, to B, decomposed in E):
	pb := nvector.Vector{X: 0.71, Y: -0.72, Z: 0.1}.Scale(6371e3)

	// Find the geodetic latitude, longitude and height, assuming WGS-84
	// ellipsoid.

	// SOLUTION:

	// Step 1
	//
	// We have a function that converts ECEF-vectors to n-vectors:
	b := nvector.FromECEF(pb, nvector.WGS84, nvector.ZAxisNorth)

	// Step 2
	//
	// Find latitude, longitude and height:
	gc := nvector.ToGeodeticCoordinates(b.Vector, nvector.ZAxisNorth)
	h := -b.Depth

	fmt.Printf(
		"Pos B: lat, lon = %.8f, %.8f deg, height = %.8f m\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
		h,
	)

	// Output:
	// Pos B: lat, lon = 5.68507573, -45.40066326 deg, height = 95772.10761822 m
}
