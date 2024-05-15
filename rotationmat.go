package nvector

import (
	"math"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/spatial/r3"
)

// FromRotationMat converts a rotation matrix to an n-vector.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/R_EN2n_E.m
func FromRotationMat(r *r3.Mat) r3.Vec {
	return r.MulVec(r3.Vec{X: 0, Y: 0, Z: -1})
}

// ToRotationMat converts n-vector to a rotation matrix.
//
// See: https://github.com/FFI-no/n-vector/blob/82d749a67cc9f332f48c51aa969cdc277b4199f2/nvector/n_E2R_EN.m
func ToRotationMat(nv r3.Vec, opts ...options.Option) *r3.Mat {
	o := options.New(opts)

	// CoordFrame selects correct E-axes
	nvr := o.CoordFrame.MulVec(nv)

	// N coordinate frame (North-East-Down) is defined in Table 2 in Gade (2010)

	// R_EN is constructed by the following three column vectors: The x, y and z
	// basis vectors (axes) of N, each decomposed in E.

	// Find z-axis of N (Nz):
	// z-axis of N (down) points opposite to n-vector
	zx := -nvr.X
	zy := -nvr.Y
	zz := -nvr.Z

	// Find y-axis of N (East)(remember that N is singular at Poles)
	// Equation (9) in Gade (2010):
	// Ny points perpendicular to the plane formed by n-vector and Earth's spin
	// axis
	yyDir := -nvr.Z
	yzDir := nvr.Y
	yDirNorm := math.Hypot(yyDir, yzDir)
	onPoles := math.Hypot(yyDir, yzDir) == 0
	var yy, yz float64
	if onPoles {
		// Pole position: selected y-axis direction
		yy = 1
		yz = 0
	} else {
		// outside Poles:
		yy = yyDir / yDirNorm
		yz = yzDir / yDirNorm
	}

	// Find x-axis of N (North):
	// Final axis found by right hand rule
	xx := yy*zz - yz*zy
	xy := yz * zx
	xz := -yy * zx

	// Form R_EN from the unit vectors:
	// CoordFrame selects correct E-axes
	r := r3.NewMat(nil)
	r.Mul(o.CoordFrame.T(), r3.NewMat([]float64{
		xx, 0, zx,
		xy, yy, zy,
		xz, yz, zz,
	}))

	return r
}

// ToRotationMatUsingWanderAzimuth converts an n-vector and a wander azimuth
// angle to a rotation matrix.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_E_and_wa2R_EL.m
func ToRotationMatUsingWanderAzimuth(
	nv r3.Vec,
	wa float64,
	opts ...options.Option,
) *r3.Mat {
	o := options.New(opts)

	// [latitude,longitude] = n_E2lat_long(n_E);
	lat, lon := ToLatLon(nv, opts...)

	// Longitude, -latitude, and wander azimuth are the x-y-z Euler angles (about
	// new axes) for R_EL. See also the second paragraph of Section 5.2 in Gade
	// (2010):

	// CoordFrame selects correct E-axes
	r := r3.NewMat(nil)
	r.Mul(o.CoordFrame.T(), XYZToRotationMat(lon, -lat, wa))

	return r
}
