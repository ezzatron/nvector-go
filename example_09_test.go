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
	// PROBLEM:

	// Define a path from two given positions (at the surface of a spherical
	// Earth), as the great circle that goes through the two points (assuming that
	// the two positions are not antipodal).

	// Path A is given by a1 and a2:
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

	// While path B is given by b1 and b2:
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

	// Find the position C where the two paths intersect.

	// SOLUTION:

	// A convenient way to represent a great circle is by its normal vector (i.e.
	// the normal vector to the plane containing the great circle). This normal
	// vector is simply found by taking the cross product of the two n-vectors
	// defining the great circle (path). Having the normal vectors to both paths,
	// the intersection is now simply found by taking the cross product of the two
	// normal vectors:
	cTmp := a1.Cross(a2).Cross(b1.Cross(b2))

	// Note that there will be two places where the great circles intersect, and
	// thus two solutions are found. Selecting the solution that is closest to
	// e.g. a1 can be achieved by selecting the solution that has a positive dot
	// product with a1 (or the mean position from Example 7 could be used instead
	// of a1):
	c := cTmp.Scale(math.Copysign(1, cTmp.Dot(a1)))

	// Use human-friendly outputs:
	gc := nvector.ToGeodeticCoordinates(c, nvector.ZAxisNorth)

	fmt.Printf(
		"Intersection: lat, lon = %.8f, %.8f deg\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
	)

	// Output:
	// Intersection: lat, lon = 74.16344802, 180.00000000 deg
}
