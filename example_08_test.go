package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/coordframe"
	"gonum.org/v1/gonum/spatial/r3"
)

// Example 8: A and azimuth/distance to B
//
// Given position A and an azimuth/bearing and a (great circle) distance. Find
// the destination point B.
//
// See: https://www.ffi.no/en/research/n-vector/#example_8
func Example_n08AAndDistanceToB() {
	// Position A is given as n_EA_E:
	// Enter elements directly:
	// a := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})

	// or input as lat/long in deg:
	a := nvector.FromLatLon(nvector.Rad(80), nvector.Rad(-90))

	// The initial azimuth and great circle distance (s_AB), and Earth radius
	// (r_Earth) are also given:
	az := nvector.Rad(200)
	sab := 1000.0 // m
	re := 6371e3  // m, mean Earth radius

	// Find the destination point B, as n_EB_E ("The direct/first geodetic
	// problem" for a sphere)

	// SOLUTION:

	// Step1: Find unit vectors for north and east (see equations (9) and (10)
	// in Gade (2010):
	e := r3.Unit(
		r3.Cross(
			coordframe.ZAxisNorth().MulVecTrans(r3.Vec{X: 1, Y: 0, Z: 0}),
			a,
		),
	)
	// const k_north_E = cross(n_EA_E, k_east_E);
	n := r3.Cross(a, e)

	// Step2: Find the initial direction vector d_E:
	d := r3.Add(
		r3.Scale(math.Cos(az), n),
		r3.Scale(math.Sin(az), e),
	)

	// Step3: Find n_EB_E:
	b := r3.Add(
		r3.Scale(math.Cos(sab/re), a),
		r3.Scale(math.Sin(sab/re), d),
	)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	lat, lon := nvector.ToLatLon(b)

	fmt.Printf(
		"Destination: lat, long = %.8f, %.8f deg\n",
		nvector.Deg(lat),
		nvector.Deg(lon),
	)

	// Output:
	// Destination: lat, long = 79.99154867, -90.01769837 deg
}
