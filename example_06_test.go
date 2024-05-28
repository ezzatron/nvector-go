package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 6: Interpolated position
//
// Given the position of B at time t(0) and t(1). Find an interpolated position
// at time t(i).
//
// See: https://www.ffi.no/en/research/n-vector/#example_6
func Example_n06InterpolatedPosition() {
	// PROBLEM:

	// Given the position of B at time t0 and t1, pt0 and pt1:
	t0, t1, ti := 10.0, 20.0, 16.0
	pt0 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(89.9),
			Longitude: nvector.Radians(-150),
		},
		nvector.ZAxisNorth,
	)
	pt1 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(89.9),
			Longitude: nvector.Radians(150),
		},
		nvector.ZAxisNorth,
	)

	// Find an interpolated position at time ti, pti. All positions are given as
	// n-vectors.

	// SOLUTION:

	// Standard interpolation can be used directly with n-vectors:
	pti := pt0.Add(pt1.Sub(pt0).Scale((ti - t0) / (t1 - t0)))

	// Use human-friendly outputs:
	gc := nvector.ToGeodeticCoordinates(pti, nvector.ZAxisNorth)

	fmt.Printf(
		"Interpolated position: lat, lon = %.8f, %.8f deg\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
	)

	// Output:
	// Interpolated position: lat, lon = 89.91282200, 173.41322445 deg
}
