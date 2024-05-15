package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
	"gonum.org/v1/gonum/spatial/r3"
)

// Example 10: Cross track distance (cross track error)
//
// Given path A going through A(1) and A(2), and a point B. Find the cross track
// distance/cross track error between B and the path.
//
// See https://www.ffi.no/en/research/n-vector/#example_10
func Example_n10CrossTrackDistance() {
	// Position A1 and A2 and B are given as n_EA1_E, n_EA2_E, and n_EB_E:
	// Enter elements directly:
	// a1 := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})
	// a2 := r3.Unit(r3.Vec{X: -1, Y: -2, Z: 0})
	// b := r3.Unit(r3.Vec{X: 0, Y: -2, Z: 3})

	// or input as lat/long in deg:
	a1 := nvector.FromLatLon(nvector.Rad(0), nvector.Rad(0))
	a2 := nvector.FromLatLon(nvector.Rad(10), nvector.Rad(0))
	b := nvector.FromLatLon(nvector.Rad(1), nvector.Rad(0.1))

	re := 6371e3 // m, mean Earth radius

	// Find the cross track distance from path A to position B.

	// SOLUTION:

	// Find the unit normal to the great circle between n_EA1_E and n_EA2_E:
	c := r3.Unit(r3.Cross(a1, a2))

	// Find the great circle cross track distance: (acos(x) - pi/2 = -asin(x))
	sxt := -math.Asin(r3.Dot(c, b)) * re

	// Find the Euclidean cross track distance:
	dxt := -r3.Dot(c, b) * re

	fmt.Printf("Cross track distance = %.8f m, Euclidean = %.8f m\n", sxt, dxt)

	// Output:
	// Cross track distance = 11117.79911015 m, Euclidean = 11117.79346741 m
}
