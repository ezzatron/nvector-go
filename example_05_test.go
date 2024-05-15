package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
	"gonum.org/v1/gonum/spatial/r3"
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
	// a := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})
	// b := r3.Unit(r3.Vec{X: -1, Y: -2, Z: 0})

	// or input as lat/long in deg:
	a := nvector.FromLatLon(nvector.Rad(88), nvector.Rad(0))
	b := nvector.FromLatLon(nvector.Rad(89), nvector.Rad(-170))

	// m, mean Earth radius
	re := 6371e3

	// SOLUTION:

	// The great circle distance is given by equation (16) in Gade (2010):
	// Well conditioned for all angles:
	sab := math.Atan2(r3.Norm(r3.Cross(a, b)), r3.Dot(a, b)) * re

	// The Euclidean distance is given by:
	dab := r3.Norm(r3.Sub(b, a)) * re

	fmt.Printf(
		"Great circle distance = %.8f km, Euclidean distance = %.8f km\n",
		sab/1000,
		dab/1000,
	)

	// Output:
	// Great circle distance = 332.45644411 km, Euclidean distance = 332.41872486 km
}
