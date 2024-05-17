package nvector

import (
	"math"
)

// GeodeticCoordinates is a geodetic latitude and longitude.
//
// Latitude and Longitude are given in radians.
type GeodeticCoordinates struct {
	Latitude, Longitude float64
}

// FromGeodeticCoordinates converts geodetic coordinates to an n-vector.
//
// f is the coordinate frame in which the n-vector is decomposed.
//
// See: https://github.com/FFI-no/n-vector/blob/82d749a67cc9f332f48c51aa969cdc277b4199f2/nvector/lat_long2n_E.m
func FromGeodeticCoordinates(c GeodeticCoordinates, f Matrix) Vector {
	// Equation (3) from Gade (2010):
	cLat := math.Cos(c.Latitude)

	// f selects correct E-axes
	return Vector{
		math.Sin(c.Latitude),
		math.Sin(c.Longitude) * cLat,
		-math.Cos(c.Longitude) * cLat,
	}.Transform(f.Transpose())
}

// ToGeodeticCoordinates converts an n-vector to geodetic coordinates.
//
// f is the coordinate frame in which the n-vector is decomposed.
//
// See: https://github.com/FFI-no/n-vector/blob/82d749a67cc9f332f48c51aa969cdc277b4199f2/nvector/n_E2lat_long.m
func ToGeodeticCoordinates(v Vector, f Matrix) GeodeticCoordinates {
	v = v.Transform(f)

	// Equation (5) in Gade (2010):
	lon := math.Atan2(v.Y, -v.Z)

	// Equation (6) in Gade (2010) (Robust numerical solution)
	// vector component in the equatorial plane
	ec := math.Hypot(v.Y, v.Z)
	// Atan() could also be used since latitude is within [-pi/2,pi/2]
	lat := math.Atan2(v.X, ec)

	// latitude = asin(v.X) is a theoretical solution, but close to the poles it
	// is ill-conditioned which may lead to numerical inaccuracies (and it will
	// give imaginary results for v.Norm()>1)

	return GeodeticCoordinates{lat, lon}
}
