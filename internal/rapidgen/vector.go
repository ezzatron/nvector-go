package rapidgen

import (
	"math"

	"github.com/ezzatron/nvector-go"
	"pgregory.net/rapid"
)

// UnitVector creates a rapid generator for unit vectors.
func UnitVector() *rapid.Generator[nvector.Vector] {
	return rapid.Custom(func(t *rapid.T) nvector.Vector {
		// based on https://github.com/mrdoob/three.js/blob/a2e9ee8204b67f9dca79f48cf620a34a05aa8126/src/math/Vector3.js#L695
		// https://mathworld.wolfram.com/SpherePointPicking.html
		theta := rapid.Float64Range(0, math.Pi*2).Draw(t, "theta")
		u := rapid.Float64Range(-1, 1).Draw(t, "u")
		c := math.Sqrt(1 - u*u)

		return nvector.Vector{
			X: c * math.Cos(theta),
			Y: u,
			Z: c * math.Sin(theta),
		}
	})
}

// VectorRange creates a rapid generator for vectors with a range of magnitudes.
func VectorRange(min, max float64) *rapid.Generator[nvector.Vector] {
	return rapid.Custom(func(t *rapid.T) nvector.Vector {
		return UnitVector().Draw(t, "v").Scale(
			rapid.Float64Range(min, max).Draw(t, "magnitude"),
		)
	})
}
