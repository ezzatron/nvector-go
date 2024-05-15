package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
	"gonum.org/v1/gonum/spatial/r3"
)

// Example 9: Intersection of two paths
//
// Given path A going through A(1) and A(2), and path B going through B(1) and
// B(2). Find the intersection of the two paths.
//
// See: https://www.ffi.no/en/research/n-vector/#example_9
func Example_n09IntersectionOfPaths() {
	// Two paths A and B are given by two pairs of positions:
	// Enter elements directly:
	// a1 := r3.Unit(r3.Vec{X: 0, Y: 0, Z: 1})
	// a2 := r3.Unit(r3.Vec{X: -1, Y: 0, Z: 1})
	// b1 := r3.Unit(r3.Vec{X: -2, Y: -2, Z: 4})
	// b2 := r3.Unit(r3.Vec{X: -2, Y: 2, Z: 2})

	// or input as lat/long in deg:
	a1 := nvector.FromLatLon(nvector.Rad(50), nvector.Rad(180))
	a2 := nvector.FromLatLon(nvector.Rad(90), nvector.Rad(180))
	b1 := nvector.FromLatLon(nvector.Rad(60), nvector.Rad(160))
	b2 := nvector.FromLatLon(nvector.Rad(80), nvector.Rad(-140))

	// SOLUTION:

	// Find the intersection between the two paths, n_EC_E:
	cTmp := r3.Cross(r3.Cross(a1, a2), r3.Cross(b1, b2))

	// n_EC_E_tmp is one of two solutions, the other is -n_EC_E_tmp. Select the
	// one that is closest to n_EA1_E, by selecting sign from the dot product
	// between n_EC_E_tmp and n_EA1_E:
	c := r3.Scale(math.Copysign(1, r3.Dot(cTmp, a1)), cTmp)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	lat, lon := nvector.ToLatLon(c)

	fmt.Printf(
		"Intersection: lat, long = %.8f, %.8f deg\n",
		nvector.Deg(lat),
		nvector.Deg(lon),
	)

	// Output:
	// Intersection: lat, long = 74.16344802, 180.00000000 deg
}
