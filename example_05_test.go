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
	// PROBLEM:

	// Given two positions A and B as n-vectors:
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

	// Find the surface distance (i.e. great circle distance). The heights of A
	// and B are not relevant (i.e. if they do not have zero height, we seek the
	// distance between the points that are at the surface of the Earth, directly
	// above/below A and B). The Euclidean distance (chord length) should also be
	// found.

	// Use Earth radius r:
	r := 6371e3

	// SOLUTION:

	// Find the great circle distance:
	gcd := math.Atan2(a.Cross(b).Norm(), a.Dot(b)) * r

	// Find the Euclidean distance:
	ed := b.Sub(a).Norm() * r

	fmt.Printf(
		"Great circle distance = %.8f m, Euclidean distance = %.8f m\n",
		gcd,
		ed,
	)

	// Output:
	// Great circle distance = 332456.44410534 m, Euclidean distance = 332418.72485681 m
}
