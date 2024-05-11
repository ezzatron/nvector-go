package equality

import (
	"math"

	"gonum.org/v1/gonum/floats/scalar"
)

// EqualToRadians returns true if the two radians are equal.
func EqualToRadians(a, b, tol float64) bool {
	// Normalize the radians to the range [0, 2π).
	a = a - 2*math.Pi*math.Floor(a/(2*math.Pi))
	b = b - 2*math.Pi*math.Floor(b/(2*math.Pi))

	return scalar.EqualWithinAbs(a, b, tol)
}
