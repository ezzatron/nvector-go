package rapidgen

import (
	"math"

	"gonum.org/v1/gonum/spatial/r3"
	"pgregory.net/rapid"
)

// UnitVector creates a rapid generator for unit vectors.
func UnitVector() *rapid.Generator[r3.Vec] {
	return rapid.Custom(func(t *rapid.T) r3.Vec {
		// based on https://github.com/mrdoob/three.js/blob/a2e9ee8204b67f9dca79f48cf620a34a05aa8126/src/math/Vector3.js#L695
		// https://mathworld.wolfram.com/SpherePointPicking.html
		theta := rapid.Float64Range(0, math.Pi*2).Draw(t, "theta")
		u := rapid.Float64Range(-1, 1).Draw(t, "u")
		c := math.Sqrt(1 - u*u)

		return r3.Vec{
			X: c * math.Cos(theta),
			Y: u,
			Z: c * math.Sin(theta),
		}
	})
}

// VectorRange creates a rapid generator for vectors with a range of magnitudes.
func VectorRange(min, max float64) *rapid.Generator[r3.Vec] {
	return rapid.Custom(func(t *rapid.T) r3.Vec {
		return r3.Scale(
			rapid.Float64Range(min, max).Draw(t, "magnitude"),
			UnitVector().Draw(t, "v"),
		)
	})
}
