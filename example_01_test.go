package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
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
func Example_n01AAndBToDelta() {
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
	a := nvector.FromLatLon(nvector.Rad(aLat), nvector.Rad(aLon))
	b := nvector.FromLatLon(nvector.Rad(bLat), nvector.Rad(bLon))

	// Step2: Find p_AB_E (delta decomposed in E). WGS-84 ellipsoid is default:
	de := nvector.Delta(a, aDepth, b, bDepth)

	// Step3: Find R_EN for position A:
	r := nvector.ToRotMat(a)

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

	fmt.Printf("Delta north, east, down = %.8f, %.8f, %.8f m\n", d.X, d.Y, d.Z)
	fmt.Printf("Azimuth = %.8f deg\n", nvector.Deg(az))

	// Output:
	// Delta north, east, down = 331730.23478089, 332997.87498927, 17404.27136194 m
	// Azimuth = 45.10926324 deg
}
