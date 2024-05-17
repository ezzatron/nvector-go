package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
)

// Example 5: Surface distance
//
// Given position A and B. Find the surface distance (i.e. great circle
// distance) and the Euclidean distance.
//
// See: https://www.ffi.no/en/research/n-vector/#example_5
func Example_n05SurfaceDistance() {
	// Position A and B are given as n_EA_E and n_EB_E:
	// Enter elements directly:
	// a := nvector.Vector{X: 1, Y: 0, Z: -2}.Normalize()
	// b := nvector.Vector{X: -1, Y: -2, Z: 0}.Normalize()

	// or input as lat/long in deg:
	a := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(88),
			Longitude: nvector.Radians(0),
		},
		nvector.ZAxisNorth,
	)
	b := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(89),
			Longitude: nvector.Radians(-170),
		},
		nvector.ZAxisNorth,
	)

	// m, mean Earth radius
	r := 6371e3

	// SOLUTION:

	// The great circle distance is given by equation (16) in Gade (2010):
	// Well conditioned for all angles:
	gcd := math.Atan2(a.Cross(b).Norm(), a.Dot(b)) * r

	// The Euclidean distance is given by:
	ed := b.Sub(a).Norm() * r

	fmt.Printf(
		"Great circle distance = %.8f km, Euclidean distance = %.8f km\n",
		gcd/1000,
		ed/1000,
	)

	// Output:
	// Great circle distance = 332.45644411 km, Euclidean distance = 332.41872486 km
}
