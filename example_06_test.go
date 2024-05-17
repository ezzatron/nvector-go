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
	// Position B is given at time t0 as n_EB_E_t0 and at time t1 as n_EB_E_t1:
	// Enter elements directly:
	// pt0 := nvector.Vector{X: 1, Y: 0, Z: -2}.Normalize()
	// pt1 := nvector.Vector{X: -1, Y: -2, Z: 0}.Normalize()

	// or input as lat/long in deg:
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

	// The times are given as:
	t0 := 10.0
	t1 := 20.0
	ti := 16.0 // time of interpolation

	// Find the interpolated position at time ti, n_EB_E_ti

	// SOLUTION:

	// Using standard interpolation:
	pti := pt0.Add(pt1.Sub(pt0).Scale((ti - t0) / (t1 - t0)))

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	gc := nvector.ToGeodeticCoordinates(pti, nvector.ZAxisNorth)

	fmt.Printf(
		"Interpolated position: lat, long = %.8f, %.8f deg\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
	)

	// Output:
	// Interpolated position: lat, long = 89.91282200, 173.41322445 deg
}
