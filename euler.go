package nvector

import (
	"math"

	"gonum.org/v1/gonum/spatial/r3"
)

// eps is the machine epsilon.
var eps = math.Nextafter(1, 2) - 1

// nSingularityEps is the number of machine epsilons to consider for an Euler
// singularity.
var nSingularityEps = 10.0

// XYZToRotationMat converts Euler angles in XYZ order to a rotation matrix.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/xyz2R.m
func XYZToRotationMat(x, y, z float64) *r3.Mat {
	cz := math.Cos(z)
	sz := math.Sin(z)
	cy := math.Cos(y)
	sy := math.Sin(y)
	cx := math.Cos(x)
	sx := math.Sin(x)

	return r3.NewMat([]float64{
		cy * cz, -cy * sz, sy,
		sy*sx*cz + cx*sz, -sy*sx*sz + cx*cz, -cy * sx,
		-sy*cx*cz + sx*sz, sy*cx*sz + sx*cz, cy * cx,
	})
}

// ZYXToRotationMat converts Euler angles in ZYX order to a rotation matrix.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/zyx2R.m
func ZYXToRotationMat(z, y, x float64) *r3.Mat {
	cz := math.Cos(z)
	sz := math.Sin(z)
	cy := math.Cos(y)
	sy := math.Sin(y)
	cx := math.Cos(x)
	sx := math.Sin(x)

	return r3.NewMat([]float64{
		cz * cy, -sz*cx + cz*sy*sx, sz*sx + cz*sy*cx,
		sz * cy, cz*cx + sz*sy*sx, -cz*sx + sz*sy*cx,
		-sy, cy * sx, cy * cx,
	})
}

// RotationMatToXYZ converts a rotation matrix to Euler angles in XYZ order.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/R2xyz.m
func RotationMatToXYZ(r *r3.Mat) (x, y, z float64) {
	// cy is based on as many elements as possible, to average out numerical
	// errors. It is selected as the positive square root since y: [-pi/2 pi/2]
	cy := math.Sqrt((math.Pow(r.At(0, 0), 2) +
		math.Pow(r.At(0, 1), 2) +
		math.Pow(r.At(1, 2), 2) +
		math.Pow(r.At(2, 2), 2)) / 2,
	)

	// Check if (close to) Euler angle singularity:
	if cy > nSingularityEps*eps {
		// Outside singularity:
		// atan2: [-pi pi]
		z = math.Atan2(-r.At(0, 1), r.At(0, 0))
		x = math.Atan2(-r.At(1, 2), r.At(2, 2))

		sy := r.At(0, 2)

		y = math.Atan2(sy, cy)
	} else {
		// In singularity (or close to), i.e. y = +pi/2 or -pi/2:
		// Selecting y = +-pi/2, with correct sign
		y = math.Copysign(math.Pi/2, r.At(0, 2))

		// Only the sum/difference of x and z is now given, choosing x = 0:
		x = 0

		// Lower left 2x2 elements of R_AB now only consists of sin_z and cos_z.
		// Using the two whose signs are the same for both singularities:
		z = math.Atan2(r.At(1, 0), r.At(1, 1))
	}

	return x, y, z
}

// RotationMatToZYX converts a rotation matrix to Euler angles in ZYX order.
//
// See: https://github.com/FFI-no/n-vector/blob/f77f43d18ddb6b8ea4e1a8bb23a53700af965abb/nvector/R2zyx.m
func RotationMatToZYX(r *r3.Mat) (z, y, x float64) {
	// cy is based on as many elements as possible, to average out numerical
	// errors. It is selected as the positive square root since y: [-pi/2 pi/2]
	cy := math.Sqrt((math.Pow(r.At(0, 0), 2) +
		math.Pow(r.At(1, 0), 2) +
		math.Pow(r.At(2, 1), 2) +
		math.Pow(r.At(2, 2), 2)) / 2,
	)

	// Check if (close to) Euler angle singularity:
	if cy > nSingularityEps*eps {
		// Outside singularity:
		// atan2: [-pi pi]
		z = math.Atan2(r.At(1, 0), r.At(0, 0))
		x = math.Atan2(r.At(2, 1), r.At(2, 2))

		sy := -r.At(2, 0)

		y = math.Atan2(sy, cy)
	} else {
		// In singularity (or close to), i.e. y = +pi/2 or -pi/2:
		// Selecting y = +-pi/2, with correct sign
		y = math.Copysign(math.Pi/2, r.At(2, 0))

		// Only the sum/difference of x and z is now given, choosing x = 0:
		x = 0

		// Upper right 2x2 elements of R_AB now only consists of sin_z and cos_z.
		// Using the two whose signs are the same for both singularities:
		z = math.Atan2(-r.At(0, 1), r.At(1, 1))
	}

	return z, y, x
}
