package nvector

import (
	"math"
)

// FromRotationMatrix converts a rotation matrix to an n-vector.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/R_EN2n_E.m
func FromRotationMatrix(r Matrix) Vector {
	return Vector{0, 0, -1}.Transform(r)
}

// ToRotationMatrix converts n-vector to a rotation matrix.
//
// f is the coordinate frame in which the n-vector is decomposed.
//
// See: https://github.com/FFI-no/n-vector/blob/82d749a67cc9f332f48c51aa969cdc277b4199f2/nvector/n_E2R_EN.m
func ToRotationMatrix(v Vector, f Matrix) Matrix {
	// f selects correct E-axes
	v = v.Transform(f)

	// N coordinate frame (North-East-Down) is defined in Table 2 in Gade (2010)

	// R_EN is constructed by the following three column vectors: The x, y and z
	// basis vectors (axes) of N, each decomposed in E.

	// Find z-axis of N (Nz):
	// z-axis of N (down) points opposite to n-vector
	zx := -v.X
	zy := -v.Y
	zz := -v.Z

	// Find y-axis of N (East)(remember that N is singular at Poles)
	// Equation (9) in Gade (2010):
	// Ny points perpendicular to the plane formed by n-vector and Earth's spin
	// axis
	yyDir := -v.Z
	yzDir := v.Y
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
	// f selects correct E-axes
	return f.Transpose().Multiply(Matrix{
		xx, 0, zx,
		xy, yy, zy,
		xz, yz, zz,
	})
}

// ToRotationMatrixUsingWanderAzimuth converts an n-vector and a wander azimuth
// angle to a rotation matrix.
//
// w is the wander azimuth angle in radians. f is the coordinate frame in which
// the n-vector is decomposed.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/n_E_and_wa2R_EL.m
func ToRotationMatrixUsingWanderAzimuth(
	v Vector,
	w float64,
	f Matrix,
) Matrix {
	l := ToGeodeticCoordinates(v, f)

	// Longitude, -latitude, and wander azimuth are the x-y-z Euler angles (about
	// new axes) for R_EL. See also the second paragraph of Section 5.2 in Gade
	// (2010):

	// f selects correct E-axes
	return f.Transpose().Multiply(
		EulerXYZToRotationMatrix(EulerXYZ{l.Longitude, -l.Latitude, w}),
	)
}
