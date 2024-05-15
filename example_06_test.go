package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
	"gonum.org/v1/gonum/spatial/r3"
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
	// pt0 := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})
	// pt1 := r3.Unit(r3.Vec{X: -1, Y: -2, Z: 0})

	// or input as lat/long in deg:
	pt0 := nvector.FromLatLon(nvector.Rad(89.9), nvector.Rad(-150))
	pt1 := nvector.FromLatLon(nvector.Rad(89.9), nvector.Rad(150))

	// The times are given as:
	t0 := 10.0
	t1 := 20.0
	ti := 16.0 // time of interpolation

	// Find the interpolated position at time ti, n_EB_E_ti

	// SOLUTION:

	// Using standard interpolation:
	pti := r3.Unit(
		r3.Add(
			pt0,
			r3.Scale(
				(ti-t0)/(t1-t0),
				r3.Sub(pt1, pt0),
			),
		),
	)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	lat, lon := nvector.ToLatLon(pti)

	fmt.Printf(
		"Interpolated position: lat, long = %.8f, %.8f deg\n",
		nvector.Deg(lat),
		nvector.Deg(lon),
	)

	// Output:
	// Interpolated position: lat, long = 89.91282200, 173.41322445 deg
}
