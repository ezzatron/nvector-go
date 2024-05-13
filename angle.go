package nvector

import "math"

// Deg converts an angle in radians to degrees.
func Deg(r float64) float64 {
	return r * 180 / math.Pi
}

// Rad converts an angle in degrees to radians.
func Rad(d float64) float64 {
	return d * math.Pi / 180
}
