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
	// PROBLEM:

	// Path A is given by the two n-vectors a1 and a2 (as in the previous
	// example):
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

	// And a position B is given by b:
	b := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(1),
			Longitude: nvector.Radians(0.1),
		},
		nvector.ZAxisNorth,
	)

	// Find the cross track distance between the path A (i.e. the great circle
	// through a1 and a2) and the position B (i.e. the shortest distance at the
	// surface, between the great circle and B). Also, find the Euclidean distance
	// between B and the plane defined by the great circle.

	// Use Earth radius r:
	r := 6371e3

	// SOLUTION:

	// First, find the normal to the great circle, with direction given by the
	// right hand rule and the direction of travel:
	c := a1.Cross(a2).Normalize()

	// Find the great circle cross track distance:
	gcd := -math.Asin(c.Dot(b)) * r

	// Finding the Euclidean distance is even simpler, since it is the projection
	// of b onto c, thus simply the dot product:
	ed := -c.Dot(b) * r

	// For both gcd and ed, positive answers means that B is to the right of the
	// track.

	fmt.Printf("Cross track distance = %.8f m, Euclidean = %.8f m\n", gcd, ed)

	// Output:
	// Cross track distance = 11117.79911015 m, Euclidean = 11117.79346741 m
}
