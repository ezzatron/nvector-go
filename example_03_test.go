package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
	"gonum.org/v1/gonum/spatial/r3"
)

// Example 3: ECEF-vector to geodetic latitude
//
// Given an ECEF-vector of a position. Find geodetic latitude, longitude and
// height (using WGS-84 ellipsoid).
//
// See: https://www.ffi.no/en/research/n-vector/#example_3
func Example_n03ECEFToLatLon() {
	// Position B is given as p_EB_E ("ECEF-vector")
	pb := r3.Scale(6371e3, r3.Vec{X: 0.71, Y: -0.72, Z: 0.1}) // m

	// Find position B as geodetic latitude, longitude and height

	// SOLUTION:

	// Find n-vector from the p-vector:
	nvb, db := nvector.FromECEF(pb)

	// Convert to lat, long and height:
	lat, lon := nvector.ToLatLon(nvb)
	hb := -db

	fmt.Printf(
		"Pos B: lat, long = %.8f, %.8f deg, height = %.8f m\n",
		nvector.Deg(lat),
		nvector.Deg(lon),
		hb,
	)

	// Output:
	// Pos B: lat, long = 5.68507573, -45.40066326 deg, height = 95772.10761822 m
}
