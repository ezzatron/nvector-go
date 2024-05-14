package nvector_test

import (
	"fmt"
	"math"

	. "github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/coordframe"
	"github.com/ezzatron/nvector-go/ellipsoid"
	"gonum.org/v1/gonum/spatial/r3"
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
func Example_n01() {
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

	// Step1: Convert to n-vectors (rad() converts to radians):
	a := FromLatLon(Rad(aLat), Rad(aLon))
	b := FromLatLon(Rad(bLat), Rad(bLon))

	// Step2: Find p_AB_E (delta decomposed in E). WGS-84 ellipsoid is default:
	de := ToDelta(a, aDepth, b, bDepth)

	// Step3: Find R_EN for position A:
	r := ToRotMat(a)

	// Step4: Find p_AB_N
	d := r.MulVecTrans(de)
	// (Note the transpose of R_EN: The "closest-rule" says that when decomposing,
	// the frame in the subscript of the rotation matrix that is closest to the
	// vector, should equal the frame where the vector is decomposed. Thus the
	// calculation R_NE*p_AB_E is correct, since the vector is decomposed in E,
	// and E is closest to the vector. In the above example we only had R_EN, and
	// thus we must transpose it: R_EN' = R_NE)

	// Step5: Also find the direction (azimuth) to B, relative to north:
	az := math.Atan2(d.Y, d.X)

	fmt.Printf("Delta north, east, down = %.8f, %.8f, %.8f m\n", d.X, d.Y, d.Z)
	fmt.Printf("Azimuth = %.8f deg\n", Deg(az))

	// Output:
	// Delta north, east, down = 331730.23478089, 332997.87498927, 17404.27136194 m
	// Azimuth = 45.10926324 deg
}

// Example 2: B and delta to C
//
// Given the position of vehicle B and a bearing and distance to an object C.
// Find the exact position of C. Use WGS-72 ellipsoid.
//
// See: https://www.ffi.no/en/research/n-vector/#example_2
func Example_n02() {
	// delta vector from B to C, decomposed in B is given:
	bc := r3.Vec{X: 3000, Y: 2000, Z: 100}

	// Position and orientation of B is given:
	// unit to get unit length of vector
	b := r3.Unit(r3.Vec{X: 1, Y: 2, Z: 3})
	bDepth := -400.0
	// the three angles are yaw, pitch, and roll
	r := EulerZYXToRotMat(Rad(10), Rad(20), Rad(30))

	// A custom reference ellipsoid is given (replacing WGS-84):
	// (WGS-72)
	opts := []Option{WithEllipsoid(ellipsoid.WGS72())}

	// Find the position of C.

	// SOLUTION:

	// Step1: Find R_EN:
	rb := ToRotMat(b)

	// Step2: Find R_EB, from R_EN and R_NB:
	// Note: closest frames cancel
	reb := r3.NewMat(nil)
	reb.Mul(rb, r)

	// Step3: Decompose the delta vector in E:
	// no transpose of R_EB, since the vector is in B
	bce := reb.MulVec(bc)

	// Step4: Find the position of C, using the functions that goes from one
	// position and a delta, to a new position:
	c, cDepth := FromDelta(b, bDepth, bce, opts...)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	lat, lon := ToLatLon(c)

	// Here we also assume that the user wants the output to be height (= -depth):
	cHeight := -cDepth

	fmt.Printf(
		"Pos C: lat, long = %.8f, %.8f deg, height = %.8f m\n",
		Deg(lat),
		Deg(lon),
		cHeight,
	)

	// Output:
	// Pos C: lat, long = 53.32637826, 63.46812344 deg, height = 406.00719607 m
}

// Example 3: ECEF-vector to geodetic latitude
//
// Given an ECEF-vector of a position. Find geodetic latitude, longitude and
// height (using WGS-84 ellipsoid).
//
// See: https://www.ffi.no/en/research/n-vector/#example_3
func Example_n03() {
	// Position B is given as p_EB_E ("ECEF-vector")
	pb := r3.Scale(6371e3, r3.Vec{X: 0.71, Y: -0.72, Z: 0.1}) // m

	// Find position B as geodetic latitude, longitude and height

	// SOLUTION:

	// Find n-vector from the p-vector:
	nvb, db := FromECEF(pb)

	// Convert to lat, long and height:
	lat, lon := ToLatLon(nvb)
	hb := -db

	fmt.Printf(
		"Pos B: lat, long = %.8f, %.8f deg, height = %.8f m\n",
		Deg(lat),
		Deg(lon),
		hb,
	)

	// Output:
	// Pos B: lat, long = 5.68507573, -45.40066326 deg, height = 95772.10761822 m
}

// Example 4: Geodetic latitude to ECEF-vector
//
// Given geodetic latitude, longitude and height. Find the ECEF-vector (using
// WGS-84 ellipsoid).
//
// See: https://www.ffi.no/en/research/n-vector/#example_4
func Example_n04() {
	// Position B is given with lat, long and height:
	bLat := 1.0
	bLon := 2.0
	bHeight := 3.0

	// Find the vector p_EB_E ("ECEF-vector")

	// SOLUTION:

	// Step1: Convert to n-vector:
	nvb := FromLatLon(Rad(bLat), Rad(bLon))

	// Step2: Find the ECEF-vector p_EB_E:
	pb := ToECEF(nvb, -bHeight)

	fmt.Printf("p_EB_E = [%.8f, %.8f, %.8f] m\n", pb.X, pb.Y, pb.Z)

	// Output:
	// p_EB_E = [6373290.27721828, 222560.20067474, 110568.82718179] m
}

// Example 5: Surface distance
//
// Given position A and B. Find the surface distance (i.e. great circle
// distance) and the Euclidean distance.
//
// See: https://www.ffi.no/en/research/n-vector/#example_5
func Example_n05() {
	// Position A and B are given as n_EA_E and n_EB_E:
	// Enter elements directly:
	// a := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})
	// b := r3.Unit(r3.Vec{X: -1, Y: -2, Z: 0})

	// or input as lat/long in deg:
	a := FromLatLon(Rad(88), Rad(0))
	b := FromLatLon(Rad(89), Rad(-170))

	// m, mean Earth radius
	re := 6371e3

	// SOLUTION:

	// The great circle distance is given by equation (16) in Gade (2010):
	// Well conditioned for all angles:
	sab := math.Atan2(r3.Norm(r3.Cross(a, b)), r3.Dot(a, b)) * re

	// The Euclidean distance is given by:
	dab := r3.Norm(r3.Sub(b, a)) * re

	fmt.Printf(
		"Great circle distance = %.8f km, Euclidean distance = %.8f km\n",
		sab/1000,
		dab/1000,
	)

	// Output:
	// Great circle distance = 332.45644411 km, Euclidean distance = 332.41872486 km
}

// Example 6: Interpolated position
//
// Given the position of B at time t(0) and t(1). Find an interpolated position
// at time t(i).
//
// See: https://www.ffi.no/en/research/n-vector/#example_6
func Example_n06() {
	// Position B is given at time t0 as n_EB_E_t0 and at time t1 as n_EB_E_t1:
	// Enter elements directly:
	// pt0 := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})
	// pt1 := r3.Unit(r3.Vec{X: -1, Y: -2, Z: 0})

	// or input as lat/long in deg:
	pt0 := FromLatLon(Rad(89.9), Rad(-150))
	pt1 := FromLatLon(Rad(89.9), Rad(150))

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
	lat, lon := ToLatLon(pti)

	fmt.Printf(
		"Interpolated position: lat, long = %.8f, %.8f deg\n",
		Deg(lat),
		Deg(lon),
	)

	// Output:
	// Interpolated position: lat, long = 89.91282200, 173.41322445 deg
}

// Example 7: Mean position/center
//
// Given three positions A, B, and C. Find the mean position (center/midpoint).
//
// See: https://www.ffi.no/en/research/n-vector/#example_7
func Example_n07() {
	// Three positions A, B and C are given:
	// Enter elements directly:
	// a := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})
	// b := r3.Unit(r3.Vec{X: -1, Y: -2, Z: 0})
	// c := r3.Unit(r3.Vec{X: 0, Y: -2, Z: 3})

	// or input as lat/long in degrees:
	a := FromLatLon(Rad(90), Rad(0))
	b := FromLatLon(Rad(60), Rad(10))
	c := FromLatLon(Rad(50), Rad(-20))

	// SOLUTION:

	// Find the horizontal mean position, M:
	m := r3.Unit(r3.Add(r3.Add(a, b), c))

	fmt.Printf("Mean position: [%.8f, %.8f, %.8f]\n", m.X, m.Y, m.Z)

	// Output:
	// Mean position: [0.38411717, -0.04660241, 0.92210749]
}

// Example 8: A and azimuth/distance to B
//
// Given position A and an azimuth/bearing and a (great circle) distance. Find
// the destination point B.
//
// See: https://www.ffi.no/en/research/n-vector/#example_8
func Example_n08() {
	// Position A is given as n_EA_E:
	// Enter elements directly:
	// a := r3.Unit(r3.Vec{X: 1, Y: 0, Z: -2})

	// or input as lat/long in deg:
	a := FromLatLon(Rad(80), Rad(-90))

	// The initial azimuth and great circle distance (s_AB), and Earth radius
	// (r_Earth) are also given:
	az := Rad(200)
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
	lat, lon := ToLatLon(b)

	fmt.Printf(
		"Destination: lat, long = %.8f, %.8f deg\n",
		Deg(lat),
		Deg(lon),
	)

	// Output:
	// Destination: lat, long = 79.99154867, -90.01769837 deg
}

// Example 9: Intersection of two paths
//
// Given path A going through A(1) and A(2), and path B going through B(1) and
// B(2). Find the intersection of the two paths.
//
// See: https://www.ffi.no/en/research/n-vector/#example_9
func Example_n09() {
	// Two paths A and B are given by two pairs of positions:
	// Enter elements directly:
	// a1 := r3.Unit(r3.Vec{X: 0, Y: 0, Z: 1})
	// a2 := r3.Unit(r3.Vec{X: -1, Y: 0, Z: 1})
	// b1 := r3.Unit(r3.Vec{X: -2, Y: -2, Z: 4})
	// b2 := r3.Unit(r3.Vec{X: -2, Y: 2, Z: 2})

	// or input as lat/long in deg:
	a1 := FromLatLon(Rad(50), Rad(180))
	a2 := FromLatLon(Rad(90), Rad(180))
	b1 := FromLatLon(Rad(60), Rad(160))
	b2 := FromLatLon(Rad(80), Rad(-140))

	// SOLUTION:

	// Find the intersection between the two paths, n_EC_E:
	cTmp := r3.Cross(r3.Cross(a1, a2), r3.Cross(b1, b2))

	// n_EC_E_tmp is one of two solutions, the other is -n_EC_E_tmp. Select the
	// one that is closest to n_EA1_E, by selecting sign from the dot product
	// between n_EC_E_tmp and n_EA1_E:
	c := r3.Scale(math.Copysign(1, r3.Dot(cTmp, a1)), cTmp)

	// When displaying the resulting position for humans, it is more convenient
	// to see lat, long:
	lat, lon := ToLatLon(c)

	fmt.Printf(
		"Intersection: lat, long = %.8f, %.8f deg\n",
		Deg(lat),
		Deg(lon),
	)

	// Output:
	// Intersection: lat, long = 74.16344802, 180.00000000 deg
}
