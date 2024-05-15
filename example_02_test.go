package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/ellipsoid"
	"gonum.org/v1/gonum/spatial/r3"
)

// Example 2: B and delta to C
//
// Given the position of vehicle B and a bearing and distance to an object C.
// Find the exact position of C. Use WGS-72 ellipsoid.
//
// See: https://www.ffi.no/en/research/n-vector/#example_2
func Example_n02BAndDeltaToC() {
	// delta vector from B to C, decomposed in B is given:
	bc := r3.Vec{X: 3000, Y: 2000, Z: 100}

	// Position and orientation of B is given:
	// unit to get unit length of vector
	b := r3.Unit(r3.Vec{X: 1, Y: 2, Z: 3})
	bDepth := -400.0
	// the three angles are yaw, pitch, and roll
	r := nvector.EulerZYXToRotMat(
		nvector.Rad(10),
		nvector.Rad(20),
		nvector.Rad(30),
	)

	// A custom reference ellipsoid is given (replacing WGS-84):
	// (WGS-72)
	opts := []nvector.Option{nvector.WithEllipsoid(ellipsoid.WGS72())}

	// Find the position of C.

	// SOLUTION:

	// Step1: Find R_EN:
	rb := nvector.ToRotMat(b)

	// Step2: Find R_EB, from R_EN and R_NB:
	// Note: closest frames cancel
	reb := r3.NewMat(nil)
	reb.Mul(rb, r)

	// Step3: Decompose the delta vector in E:
	// no transpose of R_EB, since the vector is in B
	bce := reb.MulVec(bc)

	// Step4: Find the position of C, using the functions that goes from one
	// position and a delta, to a new position:
	c, cDepth := nvector.FromDelta(b, bDepth, bce, opts...)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	lat, lon := nvector.ToLatLon(c)

	// Here we also assume that the user wants the output to be height (= -depth):
	cHeight := -cDepth

	fmt.Printf(
		"Pos C: lat, long = %.8f, %.8f deg, height = %.8f m\n",
		nvector.Deg(lat),
		nvector.Deg(lon),
		cHeight,
	)

	// Output:
	// Pos C: lat, long = 53.32637826, 63.46812344 deg, height = 406.00719607 m
}
