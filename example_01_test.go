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
// See: https://www.ffi.no/en/research/n-vector/#example_1
func Example_n01AAndBToDelta() {
	// PROBLEM:

	// Given two positions, A and B as latitudes, longitudes and depths (relative
	// to Earth, E):
	aLat, aLon, aDepth := 1.0, 2.0, 3.0
	bLat, bLon, bDepth := 4.0, 5.0, 6.0

	// Find the exact vector between the two positions, given in meters north,
	// east, and down, and find the direction (azimuth) to B, relative to north.
	//
	// Details:
	//
	// - Assume WGS-84 ellipsoid. The given depths are from the ellipsoid surface.
	// - Use position A to define north, east, and down directions. (Due to the
	//   curvature of Earth and different directions to the North Pole, the north,
	//   east, and down directions will change (relative to Earth) for different
	//   places. Position A must be outside the poles for the north and east
	//   directions to be defined.

	// SOLUTION:

	// Step 1
	//
	// First, the given latitudes and longitudes are converted to n-vectors:
	a := nvector.Position{
		Vector: nvector.FromGeodeticCoordinates(
			nvector.GeodeticCoordinates{
				Latitude:  nvector.Radians(aLat),
				Longitude: nvector.Radians(aLon),
			},
			nvector.ZAxisNorth,
		),
		Depth: aDepth,
	}
	b := nvector.Position{
		Vector: nvector.FromGeodeticCoordinates(
			nvector.GeodeticCoordinates{
				Latitude:  nvector.Radians(bLat),
				Longitude: nvector.Radians(bLon),
			},
			nvector.ZAxisNorth,
		),
		Depth: bDepth,
	}

	// Step 2
	//
	// When the positions are given as n-vectors (and depths), it is easy to find
	// the delta vector decomposed in E. No ellipsoid is specified when calling
	// the function, thus WGS-84 (default) is used:
	abE := nvector.Delta(a, b, nvector.WGS84, nvector.ZAxisNorth)

	// Step 3
	//
	// We now have the delta vector from A to B, but the three coordinates of the
	// vector are along the Earth coordinate frame E, while we need the
	// coordinates to be north, east and down. To get this, we define a
	// North-East-Down coordinate frame called N, and then we need the rotation
	// matrix (direction cosine matrix) rEN to go between E and N. We have a
	// simple function that calculates rEN from an n-vector, and we use this
	// function (using the n-vector at position A):
	rEN := nvector.ToRotationMatrix(a.Vector, nvector.ZAxisNorth)

	// Step 4
	//
	// Now the delta vector is easily decomposed in N. Since the vector is
	// decomposed in E, we must use rNE (rNE is the transpose of rEN):
	abN := abE.Transform(rEN.Transpose())

	// Step 5
	//
	// The three components of abN are the north, east and down displacements from
	// A to B in meters. The azimuth is simply found from element 1 and 2 of the
	// vector (the north and east components):
	azimuth := math.Atan2(abN.Y, abN.X)

	fmt.Printf("Delta north, east, down = %.8f, %.8f, %.8f m\n", abN.X, abN.Y, abN.Z)
	fmt.Printf("Azimuth = %.8f deg\n", nvector.Degrees(azimuth))

	// Output:
	// Delta north, east, down = 331730.23478089, 332997.87498927, 17404.27136194 m
	// Azimuth = 45.10926324 deg
}
