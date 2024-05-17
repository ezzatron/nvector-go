package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
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
	// a := nvector.Vector{X: 1, Y: 0, Z: -2}.Normalize()

	// or input as lat/long in deg:
	a := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(80),
			Longitude: nvector.Radians(-90),
		},
		nvector.ZAxisNorth,
	)

	// The initial azimuth and great circle distance (s_AB), and Earth radius
	// (r_Earth) are also given:
	az := nvector.Radians(200)
	gcd := 1000.0 // m
	r := 6371e3   // m, mean Earth radius

	// Find the destination point B, as n_EB_E ("The direct/first geodetic
	// problem" for a sphere)

	// SOLUTION:

	// Step1: Find unit vectors for north and east (see equations (9) and (10)
	// in Gade (2010):
	e := nvector.Vector{X: 1, Y: 0, Z: 0}.
		Transform(nvector.ZAxisNorth.Transpose()).
		Cross(a).
		Normalize()
	n := a.Cross(e)

	// Step2: Find the initial direction vector d_E:
	d := n.Scale(math.Cos(az)).Add(e.Scale(math.Sin(az)))

	// Step3: Find n_EB_E:
	b := a.Scale(math.Cos(gcd / r)).Add(d.Scale(math.Sin(gcd / r)))

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	gc := nvector.ToGeodeticCoordinates(b, nvector.ZAxisNorth)

	fmt.Printf(
		"Destination: lat, long = %.8f, %.8f deg\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
	)

	// Output:
	// Destination: lat, long = 79.99154867, -90.01769837 deg
}
