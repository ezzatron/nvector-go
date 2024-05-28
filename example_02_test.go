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
	// PROBLEM:

	// A radar or sonar attached to a vehicle B (Body coordinate frame) measures
	// the distance and direction to an object C. We assume that the distance and
	// two angles measured by the sensor (typically bearing and elevation relative
	// to B) are already converted (by converting from spherical to Cartesian
	// coordinates) to the vector bcB (i.e. the vector from B to C, decomposed in
	// B):
	bcB := nvector.Vector{X: 3000, Y: 2000, Z: 100}

	// The position of B is given as an n-vector and a depth:
	b := nvector.Position{
		Vector: nvector.Vector{X: 1, Y: 2, Z: 3}.Normalize(),
		Depth:  -400.0,
	}

	// The orientation (attitude) of B is given as rNB, specified as yaw, pitch,
	// roll:
	rNB := nvector.EulerZYXToRotationMatrix(nvector.EulerZYX{
		Z: nvector.Radians(10),
		Y: nvector.Radians(20),
		X: nvector.Radians(30),
	})

	// Use the WGS-72 ellipsoid:
	e := nvector.WGS72

	// Find the exact position of object C as an n-vector and a depth.

	// SOLUTION:

	// Step 1
	//
	// The delta vector is given in B. It should be decomposed in E before using
	// it, and thus we need rEB. This matrix is found from the matrices rEN and
	// rNB, and we need to find rEN, as in Example 1:
	rEN := nvector.ToRotationMatrix(b.Vector, nvector.ZAxisNorth)

	// Step 2
	//
	// Now, we can find rEB y using that the closest frames cancel when
	// multiplying two rotation matrices (i.e. N is cancelled here):
	rEB := rEN.Multiply(rNB)

	// Step 3
	//
	// The delta vector is now decomposed in E:
	bcE := bcB.Transform(rEB)

	// Step 4
	//
	// It is now easy to find the position of C using destination (with custom
	// ellipsoid overriding the default WGS-84):
	c := nvector.Destination(b, bcE, e, nvector.ZAxisNorth)

	// Use human-friendly outputs:
	gc := nvector.ToGeodeticCoordinates(c.Vector, nvector.ZAxisNorth)
	h := -c.Depth

	fmt.Printf(
		"Pos C: lat, lon = %.8f, %.8f deg, height = %.8f m\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
		h,
	)

	// Output:
	// Pos C: lat, lon = 53.32637826, 63.46812344 deg, height = 406.00719607 m
}
