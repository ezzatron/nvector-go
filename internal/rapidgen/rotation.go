package rapidgen

import (
	"math"

	"gonum.org/v1/gonum/num/quat"
	"gonum.org/v1/gonum/spatial/r3"
	"pgregory.net/rapid"
)

// Quaternion creates a rapid generator for quaternions.
func Quaternion() *rapid.Generator[quat.Number] {
	return rapid.Custom(func(t *rapid.T) quat.Number {
		// based on https://github.com/mrdoob/three.js/blob/a2e9ee8204b67f9dca79f48cf620a34a05aa8126/src/math/Quaternion.js#L592
		// Ken Shoemake
		// Uniform random rotations
		// D. Kirk, editor, Graphics Gems III, pages 124-132. Academic Press, New York, 1992.
		theta1 := rapid.Float64Range(0, math.Pi*2).Draw(t, "theta1")
		theta2 := rapid.Float64Range(0, math.Pi*2).Draw(t, "theta2")
		x0 := rapid.Float64Range(0, 1).Draw(t, "x0")

		r1 := math.Sqrt(1 - x0)
		r2 := math.Sqrt(x0)

		x := r1 * math.Sin(theta1)
		y := r1 * math.Cos(theta1)
		z := r2 * math.Sin(theta2)
		w := r2 * math.Cos(theta2)

		return quat.Number{Real: w, Imag: x, Jmag: y, Kmag: z}
	})
}

// RotationMat creates a rapid generator for rotation matrices.
func RotationMat() *rapid.Generator[*r3.Mat] {
	return rapid.Custom(func(t *rapid.T) *r3.Mat {
		// based on https://github.com/rawify/Quaternion.js/blob/c3834673b502e64e1866dbbf13568c0be93e52cc/q.js#L791
		q := Quaternion().Draw(t, "quaternion")
		w, x, y, z := q.Real, q.Imag, q.Jmag, q.Kmag

		wx := w * x
		wy := w * y
		wz := w * z
		xx := x * x
		xy := x * y
		xz := x * z
		yy := y * y
		yz := y * z
		zz := z * z

		return r3.NewMat([]float64{
			1 - 2*(yy+zz), 2 * (xy - wz), 2 * (xz + wy),
			2 * (xy + wz), 1 - 2*(xx+zz), 2 * (yz - wx),
			2 * (xz - wy), 2 * (yz + wx), 1 - 2*(xx+yy),
		})
	})
}
