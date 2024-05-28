package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 7: Mean position/center
//
// Given three positions A, B, and C. Find the mean position (center/midpoint).
//
// See: https://www.ffi.no/en/research/n-vector/#example_7
func Example_n07MeanPosition() {
	// PROBLEM:

	// Three positions A, B, and C are given as n-vectors:
	a := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(90),
			Longitude: nvector.Radians(0),
		},
		nvector.ZAxisNorth,
	)
	b := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(60),
			Longitude: nvector.Radians(10),
		},
		nvector.ZAxisNorth,
	)
	c := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(50),
			Longitude: nvector.Radians(-20),
		},
		nvector.ZAxisNorth,
	)

	// Find the mean position, M. Note that the calculation is independent of the
	// heights/depths of the positions.

	// SOLUTION:

	// The mean position is simply given by the mean n-vector:
	m := a.Add(b).Add(c).Normalize()

	fmt.Printf("Mean position: [%.8f, %.8f, %.8f]\n", m.X, m.Y, m.Z)

	// Output:
	// Mean position: [0.38411717, -0.04660241, 0.92210749]
}
