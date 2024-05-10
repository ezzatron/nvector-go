package rapidgen

import "pgregory.net/rapid"

// Latitude creates a rapid generator for latitudes.
func Latitude() *rapid.Generator[float64] {
	return rapid.Float64Range(-90, 90)
}

// Longitude creates a rapid generator for longitudes.
func Longitude() *rapid.Generator[float64] {
	return rapid.Float64Range(-180, 180)
}
