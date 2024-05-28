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
	// PROBLEM:

	// Position A is given as n-vector:
	a := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(80),
			Longitude: nvector.Radians(-90),
		},
		nvector.ZAxisNorth,
	)

	// We also have an initial direction of travel given as an azimuth (bearing)
	// relative to north (clockwise), and finally the distance to travel along a
	// great circle is given:
	azimuth := nvector.Radians(200)
	gcd := 1000.0

	// Use Earth radius r:
	r := 6371e3

	// Find the destination point B.
	//
	// In geodesy, this is known as "The first geodetic problem" or "The direct
	// geodetic problem" for a sphere, and we see that this is similar to Example
	// 2, but now the delta is given as an azimuth and a great circle distance.
	// "The second/inverse geodetic problem" for a sphere is already solved in
	// Examples 1 and 5.

	// SOLUTION:

	// The azimuth (relative to north) is a singular quantity (undefined at the
	// Poles), but from this angle we can find a (non-singular) quantity that is
	// more convenient when working with vector algebra: a vector d that points in
	// the initial direction. We find this from azimuth by first finding the north
	// and east vectors at the start point, with unit lengths.
	//
	// Here we have assumed that our coordinate frame E has its z-axis along the
	// rotational axis of the Earth, pointing towards the North Pole. Hence, this
	// axis is given by [1, 0, 0]:
	e := nvector.Vector{X: 1, Y: 0, Z: 0}.
		Transform(nvector.ZAxisNorth.Transpose()).
		Cross(a).
		Normalize()
	n := a.Cross(e)

	// The two vectors n and e are horizontal, orthogonal, and span the tangent
	// plane at the initial position. A unit vector d in the direction of the
	// azimuth is now given by:
	d := n.Scale(math.Cos(azimuth)).Add(e.Scale(math.Sin(azimuth)))

	// With the initial direction given as d instead of azimuth, it is now quite
	// simple to find b. We know that d and a are orthogonal, and they will span
	// the plane where b will lie. Thus, we can use sin and cos in the same manner
	// as above, with the angle traveled given by gcd / r:
	b := a.Scale(math.Cos(gcd / r)).Add(d.Scale(math.Sin(gcd / r)))

	// Use human-friendly outputs:
	gc := nvector.ToGeodeticCoordinates(b, nvector.ZAxisNorth)

	fmt.Printf(
		"Destination: lat, lon = %.8f, %.8f deg\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
	)

	// Output:
	// Destination: lat, lon = 79.99154867, -90.01769837 deg
}
