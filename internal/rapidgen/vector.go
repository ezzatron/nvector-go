package rapidgen

import (
	"math"

	"gonum.org/v1/gonum/mat"
	"pgregory.net/rapid"
)

// UnitVector creates a rapid generator for unit vectors.
func UnitVector() *rapid.Generator[mat.Vector] {
	return rapid.Custom(func(t *rapid.T) mat.Vector {
		// based on https://github.com/mrdoob/three.js/blob/a2e9ee8204b67f9dca79f48cf620a34a05aa8126/src/math/Vector3.js#L695
		// https://mathworld.wolfram.com/SpherePointPicking.html
		theta := rapid.Float64Range(0, math.Pi*2).Draw(t, "theta")
		u := rapid.Float64Range(-1, 1).Draw(t, "u")
		c := math.Sqrt(1 - u*u)

		return mat.NewVecDense(3, []float64{
			c * math.Cos(theta),
			u,
			c * math.Sin(theta),
		})
	})
}
