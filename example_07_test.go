package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
	"gonum.org/v1/gonum/spatial/r3"
)

// Example 7: Mean position/center
//
// Given three positions A, B, and C. Find the mean position (center/midpoint).
//
// See: https://www.ffi.no/en/research/n-vector/#example_7
func Example_n07MeanPosition() {
	// Three positions A, B and C are given:
	// Enter elements directly:
	// a := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})
	// b := r3.Unit(r3.Vec{X: -1, Y: -2, Z: 0})
	// c := r3.Unit(r3.Vec{X: 0, Y: -2, Z: 3})

	// or input as lat/long in degrees:
	a := nvector.FromLatLon(nvector.Rad(90), nvector.Rad(0))
	b := nvector.FromLatLon(nvector.Rad(60), nvector.Rad(10))
	c := nvector.FromLatLon(nvector.Rad(50), nvector.Rad(-20))

	// SOLUTION:

	// Find the horizontal mean position, M:
	m := r3.Unit(r3.Add(r3.Add(a, b), c))

	fmt.Printf("Mean position: [%.8f, %.8f, %.8f]\n", m.X, m.Y, m.Z)

	// Output:
	// Mean position: [0.38411717, -0.04660241, 0.92210749]
}
