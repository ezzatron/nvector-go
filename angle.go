package nvector

import "math"

// Degrees converts an angle in radians to degrees.
func Degrees(r float64) float64 {
	return r * 180 / math.Pi
}

// Radians converts an angle in degrees to radians.
func Radians(d float64) float64 {
	return d * math.Pi / 180
}
