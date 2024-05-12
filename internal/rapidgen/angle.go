package rapidgen

import (
	"math"

	"pgregory.net/rapid"
)

// Radians creates a rapid generator for radians.
func Radians() *rapid.Generator[float64] {
	return rapid.Float64Range(-math.Pi, math.Pi)
}
