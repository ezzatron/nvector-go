package nvector

import (
	"math"
)

// eulerThreshold is a small number used to avoid Euler angle singularities.
// This number was chosen by calculating (math.Nextafter(1, 2) - 1).
const eulerThreshold = 2.220446049250313e-16 * 10

// EulerXYZ is a set of Euler angles in XYZ order.
type EulerXYZ struct {
	X, Y, Z float64
}

// EulerZYX is a set of Euler angles in ZYX order.
type EulerZYX struct {
	Z, Y, X float64
}

// RotationMatrixToEulerXYZ converts a rotation matrix to Euler angles in XYZ order.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/R2xyz.m
func RotationMatrixToEulerXYZ(r Matrix) EulerXYZ {
	// cy is based on as many elements as possible, to average out numerical
	// errors. It is selected as the positive square root since y: [-pi/2 pi/2]
	cy := math.Sqrt((math.Pow(r.XX, 2) +
		math.Pow(r.XY, 2) +
		math.Pow(r.YZ, 2) +
		math.Pow(r.ZZ, 2)) / 2,
	)

	var a EulerXYZ

	// Check if (close to) Euler angle singularity:
	if cy > eulerThreshold {
		// Outside singularity:
		// atan2: [-pi pi]
		a.Z = math.Atan2(-r.XY, r.XX)
		a.X = math.Atan2(-r.YZ, r.ZZ)

		sy := r.XZ

		a.Y = math.Atan2(sy, cy)
	} else {
		// In singularity (or close to), i.e. y = +pi/2 or -pi/2:
		// Selecting y = +-pi/2, with correct sign
		a.Y = math.Copysign(math.Pi/2, r.XZ)

		// Only the sum/difference of x and z is now given, choosing x = 0:
		a.X = 0

		// Lower left 2x2 elements of R_AB now only consists of sin_z and cos_z.
		// Using the two whose signs are the same for both singularities:
		a.Z = math.Atan2(r.YX, r.YY)
	}

	return a
}

// RotationMatrixToEulerZYX converts a rotation matrix to Euler angles in ZYX order.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/R2zyx.m
func RotationMatrixToEulerZYX(r Matrix) EulerZYX {
	// cy is based on as many elements as possible, to average out numerical
	// errors. It is selected as the positive square root since y: [-pi/2 pi/2]
	cy := math.Sqrt((math.Pow(r.XX, 2) +
		math.Pow(r.YX, 2) +
		math.Pow(r.ZY, 2) +
		math.Pow(r.ZZ, 2)) / 2,
	)

	var a EulerZYX

	// Check if (close to) Euler angle singularity:
	if cy > eulerThreshold {
		// Outside singularity:
		// atan2: [-pi pi]
		a.Z = math.Atan2(r.YX, r.XX)
		a.X = math.Atan2(r.ZY, r.ZZ)

		sy := -r.ZX

		a.Y = math.Atan2(sy, cy)
	} else {
		// In singularity (or close to), i.e. y = +pi/2 or -pi/2:
		// Selecting y = +-pi/2, with correct sign
		a.Y = math.Copysign(math.Pi/2, r.ZX)

		// Only the sum/difference of x and z is now given, choosing x = 0:
		a.X = 0

		// Upper right 2x2 elements of R_AB now only consists of sin_z and cos_z.
		// Using the two whose signs are the same for both singularities:
		a.Z = math.Atan2(-r.XY, r.YY)
	}

	return a
}

// EulerXYZToRotationMatrix converts Euler angles in XYZ order to a rotation matrix.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/xyz2R.m
func EulerXYZToRotationMatrix(a EulerXYZ) Matrix {
	cz := math.Cos(a.Z)
	sz := math.Sin(a.Z)
	cy := math.Cos(a.Y)
	sy := math.Sin(a.Y)
	cx := math.Cos(a.X)
	sx := math.Sin(a.X)

	return Matrix{
		cy * cz, -cy * sz, sy,
		sy*sx*cz + cx*sz, -sy*sx*sz + cx*cz, -cy * sx,
		-sy*cx*cz + sx*sz, sy*cx*sz + sx*cz, cy * cx,
	}
}

// EulerZYXToRotationMatrix converts Euler angles in ZYX order to a rotation matrix.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/zyx2R.m
func EulerZYXToRotationMatrix(a EulerZYX) Matrix {
	cz := math.Cos(a.Z)
	sz := math.Sin(a.Z)
	cy := math.Cos(a.Y)
	sy := math.Sin(a.Y)
	cx := math.Cos(a.X)
	sx := math.Sin(a.X)

	return Matrix{
		cz * cy, -sz*cx + cz*sy*sx, sz*sx + cz*sy*cx,
		sz * cy, cz*cx + sz*sy*sx, -cz*sx + sz*sy*cx,
		-sy, cy * sx, cy * cx,
	}
}
