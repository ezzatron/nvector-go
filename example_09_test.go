package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
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
	// a1 := nvector.Vector{X: 0, Y: 0, Z: 1}.Normalize()
	// a2 := nvector.Vector{X: -1, Y: 0, Z: 1}.Normalize()
	// b1 := nvector.Vector{X: -2, Y: -2, Z: 4}.Normalize()
	// b2 := nvector.Vector{X: -2, Y: 2, Z: 2}.Normalize()

	// or input as lat/long in deg:
	a1 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(50),
			Longitude: nvector.Radians(180),
		},
		nvector.ZAxisNorth,
	)
	a2 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(90),
			Longitude: nvector.Radians(180),
		},
		nvector.ZAxisNorth,
	)
	b1 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(60),
			Longitude: nvector.Radians(160),
		},
		nvector.ZAxisNorth,
	)
	b2 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(80),
			Longitude: nvector.Radians(-140),
		},
		nvector.ZAxisNorth,
	)

	// SOLUTION:

	// Find the intersection between the two paths, n_EC_E:
	cTmp := a1.Cross(a2).Cross(b1.Cross(b2))

	// n_EC_E_tmp is one of two solutions, the other is -n_EC_E_tmp. Select the
	// one that is closest to n_EA1_E, by selecting sign from the dot product
	// between n_EC_E_tmp and n_EA1_E:
	c := cTmp.Scale(math.Copysign(1, cTmp.Dot(a1)))

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	gc := nvector.ToGeodeticCoordinates(c, nvector.ZAxisNorth)

	fmt.Printf(
		"Intersection: lat, long = %.8f, %.8f deg\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
	)

	// Output:
	// Intersection: lat, long = 74.16344802, 180.00000000 deg
}
