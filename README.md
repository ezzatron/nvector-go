# _n_-vector

_Functions for performing geographical position calculations using n-vectors_

[![Current version][badge-version-image]][badge-version-link]
[![Build status][badge-build-image]][badge-build-link]
[![Test coverage][badge-coverage-image]][badge-coverage-link]

[badge-build-image]:
  https://img.shields.io/github/actions/workflow/status/ezzatron/nvector-go/ci.yml?branch=main&style=for-the-badge
[badge-build-link]:
  https://github.com/ezzatron/nvector-go/actions/workflows/ci.yml
[badge-coverage-image]:
  https://img.shields.io/codecov/c/gh/ezzatron/nvector-go?style=for-the-badge
[badge-coverage-link]: https://codecov.io/gh/ezzatron/nvector-go
[badge-version-image]:
  https://img.shields.io/github/v/tag/ezzatron/nvector-go?include_prereleases&sort=semver&logo=go&label=github.com%2Fezzatron%2Fnvector-go&style=for-the-badge
[badge-version-link]: https://pkg.go.dev/github.com/ezzatron/nvector-go

This library is a port of the [Matlab n-vector library] by [Kenneth Gade]. All
original functions are included, although the names of the functions and
arguments have been changed in an attempt to clarify their purpose. In addition,
this library includes some extra functions for vector and matrix operations
needed to solve the [10 examples from the n-vector page].

[matlab n-vector library]: https://github.com/FFI-no/n-vector
[kenneth gade]: https://github.com/KennethGade
[10 examples from the n-vector page]: https://www.ffi.no/en/research/n-vector

See the [reference documentation] for a list of all functions and their
signatures.

[reference documentation]: https://pkg.go.dev/github.com/ezzatron/nvector-go

## Installation

```sh
go get github.com/ezzatron/nvector-go
```

## Examples

The following sections show the [10 examples from the n-vector page] implemented
using this library.

### Example 1: A and B to delta

![Illustration of example 1](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-01.png)

> Given two positions _A_ and _B_. Find the exact vector from _A_ to _B_ in
> meters north, east and down, and find the direction (azimuth/bearing) to _B_,
> relative to north. Use WGS-84 ellipsoid.
>
> https://www.ffi.no/en/research/n-vector/#example_1

```go
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
//   - Assume WGS-84 ellipsoid. The given depths are from the ellipsoid surface.
//   - Use position A to define north, east, and down directions. (Due to the
//     curvature of Earth and different directions to the North Pole, the north,
//     east, and down directions will change (relative to Earth) for different
//     places. Position A must be outside the poles for the north and east
//     directions to be defined.
//
// See: https://www.ffi.no/en/research/n-vector/#example_1
func Example_n01AAndBToDelta() {
	// Positions A and B are given in (decimal) degrees and depths:

	// Position A:
	aLat := 1.0
	aLon := 2.0
	aDepth := 3.0

	// Position B:
	bLat := 4.0
	bLon := 5.0
	bDepth := 6.0

	// Find the exact vector between the two positions, given in meters north,
	// east, and down, i.e. find p_AB_N.

	// SOLUTION:

	// Step1: Convert to n-vectors:
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

	// Step2: Find p_AB_E (delta decomposed in E):
	pe := nvector.Delta(a, b, nvector.WGS84, nvector.ZAxisNorth)

	// Step3: Find R_EN for position A:
	r := nvector.ToRotationMatrix(a.Vector, nvector.ZAxisNorth)

	// Step4: Find p_AB_N
	pn := pe.Transform(r.Transpose())
	// (Note the transpose of R_EN: The "closest-rule" says that when decomposing,
	// the frame in the subscript of the rotation matrix that is closest to the
	// vector, should equal the frame where the vector is decomposed. Thus the
	// calculation R_NE*p_AB_E is correct, since the vector is decomposed in E,
	// and E is closest to the vector. In the above example we only had R_EN, and
	// thus we must transpose it: R_EN' = R_NE)

	// Step5: Also find the direction (azimuth) to B, relative to north:
	az := math.Atan2(pn.Y, pn.X)

	fmt.Printf("Delta north, east, down = %.8f, %.8f, %.8f m\n", pn.X, pn.Y, pn.Z)
	fmt.Printf("Azimuth = %.8f deg\n", nvector.Degrees(az))

	// Output:
	// Delta north, east, down = 331730.23478089, 332997.87498927, 17404.27136194 m
	// Azimuth = 45.10926324 deg
}
```

### Example 2: B and delta to C

![Illustration of example 2](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-02.png)

> Given the position of vehicle _B_ and a bearing and distance to an object _C_.
> Find the exact position of _C_. Use WGS-72 ellipsoid.
>
> https://www.ffi.no/en/research/n-vector/#example_2

```go
package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 2: B and delta to C
//
// Given the position of vehicle B and a bearing and distance to an object C.
// Find the exact position of C. Use WGS-72 ellipsoid.
//
// See: https://www.ffi.no/en/research/n-vector/#example_2
func Example_n02BAndDeltaToC() {
	// delta vector from B to C, decomposed in B is given:
	bc := nvector.Vector{X: 3000, Y: 2000, Z: 100}

	// Position and orientation of B is given:
	// Normalize to get unit length of vector
	b := nvector.Position{
		Vector: nvector.Vector{X: 1, Y: 2, Z: 3}.Normalize(),
		Depth:  -400.0,
	}
	// the three angles are yaw, pitch, and roll
	r := nvector.EulerZYXToRotationMatrix(nvector.EulerZYX{
		Z: nvector.Radians(10),
		Y: nvector.Radians(20),
		X: nvector.Radians(30),
	})

	// A custom reference ellipsoid is given (replacing WGS-84):
	// (WGS-72)
	e := nvector.WGS72

	// Find the position of C.

	// SOLUTION:

	// Step1: Find R_EN:
	rn := nvector.ToRotationMatrix(b.Vector, nvector.ZAxisNorth)

	// Step2: Find R_EB, from R_EN and R_NB:
	// Note: closest frames cancel
	rb := rn.Multiply(r)

	// Step3: Decompose the delta vector in E:
	// no transpose of R_EB, since the vector is in B
	bce := bc.Transform(rb)

	// Step4: Find the position of C, using the functions that goes from one
	// position and a delta, to a new position:
	c := nvector.Destination(b, bce, e, nvector.ZAxisNorth)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	gc := nvector.ToGeodeticCoordinates(c.Vector, nvector.ZAxisNorth)

	// Here we also assume that the user wants the output to be height (= -depth):
	h := -c.Depth

	fmt.Printf(
		"Pos C: lat, long = %.8f, %.8f deg, height = %.8f m\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
		h,
	)

	// Output:
	// Pos C: lat, long = 53.32637826, 63.46812344 deg, height = 406.00719607 m
}
```

### Example 3: ECEF-vector to geodetic latitude

![Illustration of example 3](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-03.png)

> Given an ECEF-vector of a position. Find geodetic latitude, longitude and
> height (using WGS-84 ellipsoid).
>
> https://www.ffi.no/en/research/n-vector/#example_3

```go
package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 3: ECEF-vector to geodetic latitude
//
// Given an ECEF-vector of a position. Find geodetic latitude, longitude and
// height (using WGS-84 ellipsoid).
//
// See: https://www.ffi.no/en/research/n-vector/#example_3
func Example_n03ECEFToLatLon() {
	// Position B is given as p_EB_E ("ECEF-vector")
	b := nvector.Vector{X: 0.71, Y: -0.72, Z: 0.1}.Scale(6371e3) // m

	// Find position B as geodetic latitude, longitude and height

	// SOLUTION:

	// Find n-vector from the p-vector:
	vb := nvector.FromECEF(b, nvector.WGS84, nvector.ZAxisNorth)

	// Convert to lat, long and height:
	gc := nvector.ToGeodeticCoordinates(vb.Vector, nvector.ZAxisNorth)
	h := -vb.Depth

	fmt.Printf(
		"Pos B: lat, long = %.8f, %.8f deg, height = %.8f m\n",
		nvector.Degrees(gc.Latitude),
		nvector.Degrees(gc.Longitude),
		h,
	)

	// Output:
	// Pos B: lat, long = 5.68507573, -45.40066326 deg, height = 95772.10761822 m
}
```

### Example 4: Geodetic latitude to ECEF-vector

![Illustration of example 4](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-04.png)

> Given geodetic latitude, longitude and height. Find the ECEF-vector (using
> WGS-84 ellipsoid).
>
> https://www.ffi.no/en/research/n-vector/#example_4

```go
package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 4: Geodetic latitude to ECEF-vector
//
// Given geodetic latitude, longitude and height. Find the ECEF-vector (using
// WGS-84 ellipsoid).
//
// See: https://www.ffi.no/en/research/n-vector/#example_4
func Example_n04LatLonToECEF() {
	// Position B is given with lat, long and height:
	bLat := 1.0
	bLon := 2.0
	bHeight := 3.0

	// Find the vector p_EB_E ("ECEF-vector")

	// SOLUTION:

	// Step1: Convert to n-vector:
	b := nvector.Position{
		Vector: nvector.FromGeodeticCoordinates(
			nvector.GeodeticCoordinates{
				Latitude:  nvector.Radians(bLat),
				Longitude: nvector.Radians(bLon),
			},
			nvector.ZAxisNorth,
		),
		Depth: -bHeight,
	}

	// Step2: Find the ECEF-vector p_EB_E:
	pb := nvector.ToECEF(b, nvector.WGS84, nvector.ZAxisNorth)

	fmt.Printf("p_EB_E = [%.8f, %.8f, %.8f] m\n", pb.X, pb.Y, pb.Z)

	// Output:
	// p_EB_E = [6373290.27721828, 222560.20067474, 110568.82718179] m
}
```

### Example 5: Surface distance

![Illustration of example 5](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-05.png)

> Given position _A_ and _B_. Find the surface **distance** (i.e. great circle
> distance) and the Euclidean distance.
>
> https://www.ffi.no/en/research/n-vector/#example_5

```go
package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
)

// Example 5: Surface distance
//
// Given position A and B. Find the surface distance (i.e. great circle
// distance) and the Euclidean distance.
//
// See: https://www.ffi.no/en/research/n-vector/#example_5
func Example_n05SurfaceDistance() {
	// Position A and B are given as n_EA_E and n_EB_E:
	// Enter elements directly:
	// a := nvector.Vector{X: 1, Y: 0, Z: -2}.Normalize()
	// b := nvector.Vector{X: -1, Y: -2, Z: 0}.Normalize()

	// or input as lat/long in deg:
	a := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(88),
			Longitude: nvector.Radians(0),
		},
		nvector.ZAxisNorth,
	)
	b := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(89),
			Longitude: nvector.Radians(-170),
		},
		nvector.ZAxisNorth,
	)

	// m, mean Earth radius
	r := 6371e3

	// SOLUTION:

	// The great circle distance is given by equation (16) in Gade (2010):
	// Well conditioned for all angles:
	gcd := math.Atan2(a.Cross(b).Norm(), a.Dot(b)) * r

	// The Euclidean distance is given by:
	ed := b.Sub(a).Norm() * r

	fmt.Printf(
		"Great circle distance = %.8f km, Euclidean distance = %.8f km\n",
		gcd/1000,
		ed/1000,
	)

	// Output:
	// Great circle distance = 332.45644411 km, Euclidean distance = 332.41872486 km
}
```

### Example 6: Interpolated position

![Illustration of example 6](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-06.png)

> Given the position of _B_ at time _t<sub>0</sub>_ and _t<sub>1</sub>_. Find an
> **interpolated position** at time _t<sub>i</sub>_.
>
> https://www.ffi.no/en/research/n-vector/#example_6

```go
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
```

### Example 7: Mean position/center

![Illustration of example 7](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-07.png)

> Given three positions _A_, _B_, and _C_. Find the **mean position**
> (center/midpoint).
>
> https://www.ffi.no/en/research/n-vector/#example_7

```go
package nvector_test

import (
	"fmt"

	"github.com/ezzatron/nvector-go"
)

// Example 7: Mean position/center
//
// Given three positions A, B, and C. Find the mean position (center/midpoint).
//
// See: https://www.ffi.no/en/research/n-vector/#example_7
func Example_n07MeanPosition() {
	// Three positions A, B and C are given:
	// Enter elements directly:
	// a := nvector.Vector{X: 1, Y: 0, Z: -2}.Normalize()
	// b := nvector.Vector{X: -1, Y: -2, Z: 0}.Normalize()
	// c := nvector.Vector{X: 0, Y: -2, Z: 3}.Normalize()

	// or input as lat/long in degrees:
	a := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(90),
			Longitude: nvector.Radians(0),
		},
		nvector.ZAxisNorth,
	)
	b := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(60),
			Longitude: nvector.Radians(10),
		},
		nvector.ZAxisNorth,
	)
	c := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(50),
			Longitude: nvector.Radians(-20),
		},
		nvector.ZAxisNorth,
	)

	// SOLUTION:

	// Find the horizontal mean position, M:
	m := a.Add(b).Add(c).Normalize()

	fmt.Printf("Mean position: [%.8f, %.8f, %.8f]\n", m.X, m.Y, m.Z)

	// Output:
	// Mean position: [0.38411717, -0.04660241, 0.92210749]
}
```

### Example 8: A and azimuth/distance to B

![Illustration of example 8](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-08.png)

> Given position _A_ and an azimuth/bearing and a (great circle) distance. Find
> the **destination point** _B_.
>
> https://www.ffi.no/en/research/n-vector/#example_8

```go
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
```

### Example 9: Intersection of two paths

![Illustration of example 9](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-09.png)

> Given path _A_ going through _A<sub>1</sub>_ and _A<sub>2</sub>_, and path _B_
> going through _B<sub>1</sub>_ and _B<sub>2</sub>_. Find the **intersection**
> of the two paths.
>
> https://www.ffi.no/en/research/n-vector/#example_9

```go
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
```

### Example 10: Cross track distance (cross track error)

![Illustration of example 10](https://raw.githubusercontent.com/ezzatron/nvector-go/main/assets/img/example-10.png)

> Given path _A_ going through _A<sub>1</sub>_ and _A<sub>2</sub>_, and a point
> _B_. Find the **cross track distance**/**cross track error** between _B_ and
> the path.
>
> https://www.ffi.no/en/research/n-vector/#example_10

```go
package nvector_test

import (
	"fmt"
	"math"

	"github.com/ezzatron/nvector-go"
)

// Example 10: Cross track distance (cross track error)
//
// Given path A going through A(1) and A(2), and a point B. Find the cross track
// distance/cross track error between B and the path.
//
// See https://www.ffi.no/en/research/n-vector/#example_10
func Example_n10CrossTrackDistance() {
	// Position A1 and A2 and B are given as n_EA1_E, n_EA2_E, and n_EB_E:
	// Enter elements directly:
	// a1 := nvector.Vector{X: 1, Y: 0, Z: -2}.Normalize()
	// a2 := nvector.Vector{X: -1, Y: -2, Z: 0}.Normalize()
	// b := nvector.Vector{X: 0, Y: -2, Z: 3}.Normalize()

	// or input as lat/long in deg:
	a1 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(0),
			Longitude: nvector.Radians(0),
		},
		nvector.ZAxisNorth,
	)
	a2 := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(10),
			Longitude: nvector.Radians(0),
		},
		nvector.ZAxisNorth,
	)
	b := nvector.FromGeodeticCoordinates(
		nvector.GeodeticCoordinates{
			Latitude:  nvector.Radians(1),
			Longitude: nvector.Radians(0.1),
		},
		nvector.ZAxisNorth,
	)

	r := 6371e3 // m, mean Earth radius

	// Find the cross track distance from path A to position B.

	// SOLUTION:

	// Find the unit normal to the great circle between n_EA1_E and n_EA2_E:
	c := a1.Cross(a2).Normalize()

	// Find the great circle cross track distance: (acos(x) - pi/2 = -asin(x))
	gcd := -math.Asin(c.Dot(b)) * r

	// Find the Euclidean cross track distance:
	ed := -c.Dot(b) * r

	fmt.Printf("Cross track distance = %.8f m, Euclidean = %.8f m\n", gcd, ed)

	// Output:
	// Cross track distance = 11117.79911015 m, Euclidean = 11117.79346741 m
}
```

## Methodology

If you look at the test suite for this library, you'll see that there are very
few concrete test cases. Instead, this library uses model-based testing, powered
by [rapid], and using the [Python nvector library] as the "model", or reference
implementation.

[rapid]: https://github.com/flyingmutant/rapid
[python nvector library]: https://nvector.readthedocs.io/

In other words, this library is tested by generating large amounts of "random"
inputs, and then comparing the output with the Python library. This allowed me
to quickly port the library with a high degree of confidence in its correctness,
without a deep understanding of the underlying mathematics.

If you find any issues with the implementations, there's a good chance that the
issue will also be present in the Python library, and an equally good chance
that I won't personally understand how to fix it 😅 Still, don't let that stop
you from opening an issue or a pull request!

## References

- Gade, K. (2010). [A Non-singular Horizontal Position Representation], The
  Journal of Navigation, Volume 63, Issue 03, pp 395-417, July 2010.
- [The n-vector page]
- Ellipsoid data taken from [chrisveness/geodesy]

[a non-singular horizontal position representation]:
  https://www.navlab.net/Publications/A_Nonsingular_Horizontal_Position_Representation.pdf
[the n-vector page]: https://www.ffi.no/en/research/n-vector
[chrisveness/geodesy]: https://github.com/chrisveness/geodesy
