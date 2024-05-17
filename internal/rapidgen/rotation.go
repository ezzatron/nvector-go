package rapidgen

import (
	"math"

	"github.com/ezzatron/nvector-go"
	"pgregory.net/rapid"
)

// RotationMatrix creates a rapid generator for rotation matrices.
func RotationMatrix() *rapid.Generator[nvector.Matrix] {
	return rapid.Custom(func(t *rapid.T) nvector.Matrix {
		// based on https://github.com/rawify/Quaternion.js/blob/c3834673b502e64e1866dbbf13568c0be93e52cc/q.js#L791
		q := quaternion().Draw(t, "quaternion")
		w, x, y, z := q.W, q.X, q.Y, q.Z

		wx := w * x
		wy := w * y
		wz := w * z
		xx := x * x
		xy := x * y
		xz := x * z
		yy := y * y
		yz := y * z
		zz := z * z

		return nvector.Matrix{
			XX: 1 - 2*(yy+zz), XY: 2 * (xy - wz), XZ: 2 * (xz + wy),
			YX: 2 * (xy + wz), YY: 1 - 2*(xx+zz), YZ: 2 * (yz - wx),
			ZX: 2 * (xz - wy), ZY: 2 * (yz + wx), ZZ: 1 - 2*(xx+yy),
		}
	})
}

type quat struct {
	W, X, Y, Z float64
}

func quaternion() *rapid.Generator[quat] {
	return rapid.Custom(func(t *rapid.T) quat {
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

		return quat{w, x, y, z}
	})
}
