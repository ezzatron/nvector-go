package nvector_test

import (
	"fmt"
	"math"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/ellipsoid"
	"gonum.org/v1/gonum/spatial/r3"
)

// Example 1: A and B to delta
//
// Given two positions A and B. Find the exact vector from A to B in meters
// north, east and down, and find the direction (azimuth/bearing) to B, relative
// to north. Use WGS-84 ellipsoid.
//
//   - Assume WGS-84 ellipsoid. The given depths are from the ellipsoid surface.
//   - Use position A to define north, east, and down directions. (Due to the
//     curvature of Earth and different directions to the North Pole, the north,
//     east, and down directions will change (relative to Earth) for different
//     places. Position A must be outside the poles for the north and east
//     directions to be defined.
//
// See: https://www.ffi.no/en/research/n-vector/#example_1
func Example_n01() {
	// Positions A and B are given in (decimal) degrees and depths:

	// Position A:
	aLat := 1.0
	aLon := 2.0
	aDepth := 3.0

	// Position B:
	bLat := 4.0
	bLon := 5.0
	bDepth := 6.0

	// Find the exact vector between the two positions, given in meters north,
	// east, and down, i.e. find p_AB_N.

	// SOLUTION:

	// Step1: Convert to n-vectors (rad() converts to radians):
	a := FromLatLon(Rad(aLat), Rad(aLon))
	b := FromLatLon(Rad(bLat), Rad(bLon))

	// Step2: Find p_AB_E (delta decomposed in E). WGS-84 ellipsoid is default:
	de := ToDelta(a, aDepth, b, bDepth)

	// Step3: Find R_EN for position A:
	r := ToRotMat(a)

	// Step4: Find p_AB_N
	d := r.MulVecTrans(de)
	// (Note the transpose of R_EN: The "closest-rule" says that when decomposing,
	// the frame in the subscript of the rotation matrix that is closest to the
	// vector, should equal the frame where the vector is decomposed. Thus the
	// calculation R_NE*p_AB_E is correct, since the vector is decomposed in E,
	// and E is closest to the vector. In the above example we only had R_EN, and
	// thus we must transpose it: R_EN' = R_NE)

	// Step5: Also find the direction (azimuth) to B, relative to north:
	az := math.Atan2(d.Y, d.X)

	fmt.Printf("Delta north, east, down = %v, %v, %v m\n", d.X, d.Y, d.Z)
	fmt.Printf("Azimuth = %v deg\n", Deg(az))

	// Output:
	// Delta north, east, down = 331730.23478089436, 332997.8749892695, 17404.271361936346 m
	// Azimuth = 45.10926323826139 deg
}

// Example 2: B and delta to C
//
// Given the position of vehicle B and a bearing and distance to an object C.
// Find the exact position of C. Use WGS-72 ellipsoid.
//
// See: https://www.ffi.no/en/research/n-vector/#example_2
func Example_n02() {
	// delta vector from B to C, decomposed in B is given:
	bc := r3.Vec{X: 3000, Y: 2000, Z: 100}

	// Position and orientation of B is given:
	// unit to get unit length of vector
	b := r3.Unit(r3.Vec{X: 1, Y: 2, Z: 3})
	bDepth := -400.0
	// the three angles are yaw, pitch, and roll
	r := EulerZYXToRotMat(Rad(10), Rad(20), Rad(30))

	// A custom reference ellipsoid is given (replacing WGS-84):
	// (WGS-72)
	opts := []Option{WithEllipsoid(ellipsoid.WGS72())}

	// Find the position of C.

	// SOLUTION:

	// Step1: Find R_EN:
	rb := ToRotMat(b)

	// Step2: Find R_EB, from R_EN and R_NB:
	// Note: closest frames cancel
	reb := r3.NewMat(nil)
	reb.Mul(rb, r)

	// Step3: Decompose the delta vector in E:
	// no transpose of R_EB, since the vector is in B
	bce := reb.MulVec(bc)

	// Step4: Find the position of C, using the functions that goes from one
	// position and a delta, to a new position:
	c, cDepth := FromDelta(b, bDepth, bce, opts...)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	lat, lon := ToLatLon(c)

	// Here we also assume that the user wants the output to be height (= -depth):
	cHeight := -cDepth

	fmt.Printf("Pos C: lat, long = %v, %v deg, height = %v m\n", Deg(lat), Deg(lon), cHeight)

	// Output:
	// Pos C: lat, long = 53.32637826433106, 63.468123435147454 deg, height = 406.00719606859053 m
}
