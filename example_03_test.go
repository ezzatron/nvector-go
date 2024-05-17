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
	// Position B is given as p_EB_E ("ECEF-vector")
	b := nvector.Vector{X: 0.71, Y: -0.72, Z: 0.1}.Scale(6371e3) // m

	// Find position B as geodetic latitude, longitude and height

	// SOLUTION:

	// Find n-vector from the p-vector:
	vb := nvector.FromECEF(b, nvector.WGS84, nvector.ZAxisNorth)

	// Convert to lat, long and height:
	gc := nvector.ToGeodeticCoordinates(vb.Vector, nvector.ZAxisNorth)
	h := -vb.Depth

	fmt.Printf(
		"Pos B: lat, long = %.8f, %.8f deg, height = %.8f m\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
		h,
	)

	// Output:
	// Pos B: lat, long = 5.68507573, -45.40066326 deg, height = 95772.10761822 m
}
