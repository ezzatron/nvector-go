package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
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
	// a1 := nvector.Vector{X: 1, Y: 0, Z: -2}.Normalize()
	// a2 := nvector.Vector{X: -1, Y: -2, Z: 0}.Normalize()
	// b := nvector.Vector{X: 0, Y: -2, Z: 3}.Normalize()

	// or input as lat/long in deg:
	a1 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(0),
			Longitude: nvector.Radians(0),
		},
		nvector.ZAxisNorth,
	)
	a2 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(10),
			Longitude: nvector.Radians(0),
		},
		nvector.ZAxisNorth,
	)
	b := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(1),
			Longitude: nvector.Radians(0.1),
		},
		nvector.ZAxisNorth,
	)

	r := 6371e3 // m, mean Earth radius

	// Find the cross track distance from path A to position B.

	// SOLUTION:

	// Find the unit normal to the great circle between n_EA1_E and n_EA2_E:
	c := a1.Cross(a2).Normalize()

	// Find the great circle cross track distance: (acos(x) - pi/2 = -asin(x))
	gcd := -math.Asin(c.Dot(b)) * r

	// Find the Euclidean cross track distance:
	ed := -c.Dot(b) * r

	fmt.Printf("Cross track distance = %.8f m, Euclidean = %.8f m\n", gcd, ed)

	// Output:
	// Cross track distance = 11117.79911015 m, Euclidean = 11117.79346741 m
}
