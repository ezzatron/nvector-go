package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 2: B and delta to C
//
// Given the position of vehicle B and a bearing and distance to an object C.
// Find the exact position of C. Use WGS-72 ellipsoid.
//
// See: https://www.ffi.no/en/research/n-vector/#example_2
func Example_n02BAndDeltaToC() {
	// delta vector from B to C, decomposed in B is given:
	bc := nvector.Vector{X: 3000, Y: 2000, Z: 100}

	// Position and orientation of B is given:
	// Normalize to get unit length of vector
	b := nvector.Position{
		Vector: nvector.Vector{X: 1, Y: 2, Z: 3}.Normalize(),
		Depth:  -400.0,
	}
	// the three angles are yaw, pitch, and roll
	r := nvector.EulerZYXToRotationMatrix(nvector.EulerZYX{
		Z: nvector.Radians(10),
		Y: nvector.Radians(20),
		X: nvector.Radians(30),
	})

	// A custom reference ellipsoid is given (replacing WGS-84):
	// (WGS-72)
	e := nvector.WGS72

	// Find the position of C.

	// SOLUTION:

	// Step1: Find R_EN:
	rn := nvector.ToRotationMatrix(b.Vector, nvector.ZAxisNorth)

	// Step2: Find R_EB, from R_EN and R_NB:
	// Note: closest frames cancel
	rb := rn.Multiply(r)

	// Step3: Decompose the delta vector in E:
	// no transpose of R_EB, since the vector is in B
	bce := bc.Transform(rb)

	// Step4: Find the position of C, using the functions that goes from one
	// position and a delta, to a new position:
	c := nvector.Destination(b, bce, e, nvector.ZAxisNorth)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	gc := nvector.ToGeodeticCoordinates(c.Vector, nvector.ZAxisNorth)

	// Here we also assume that the user wants the output to be height (= -depth):
	h := -c.Depth

	fmt.Printf(
		"Pos C: lat, long = %.8f, %.8f deg, height = %.8f m\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
		h,
	)

	// Output:
	// Pos C: lat, long = 53.32637826, 63.46812344 deg, height = 406.00719607 m
}
