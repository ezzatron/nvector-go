package nvector

import (
	"math"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/spatial/r3"
)

// FromLatLon converts a geodetic latitude and longitude to an n-vector.
//
// lat and lon are given in radians. The ellipsoid has no effect on the
// calculation.
//
// See: https://github.com/FFI-no/n-vector/blob/82d749a67cc9f332f48c51aa969cdc277b4199f2/nvector/lat_long2n_E.m
func FromLatLon(lat, lon float64, opts ...Option) r3.Vec {
	o := options.New(opts)

	// Equation (3) from Gade (2010):
	cosLat := math.Cos(lat)

	// CoordFrame selects correct E-axes
	return o.CoordFrame.MulVecTrans(
		r3.Vec{
			X: math.Sin(lat),
			Y: math.Sin(lon) * cosLat,
			Z: -math.Cos(lon) * cosLat,
		},
	)
}

// ToLatLon converts an n-vector to a geodetic latitude and longitude.
//
// lat and lon are returned in radians. The ellipsoid has no effect on the
// calculation.
//
// See: https://github.com/FFI-no/n-vector/blob/82d749a67cc9f332f48c51aa969cdc277b4199f2/nvector/n_E2lat_long.m
func ToLatLon(nv r3.Vec, opts ...Option) (lat float64, lon float64) {
	o := options.New(opts)

	nvr := o.CoordFrame.MulVec(nv)

	// Equation (5) in Gade (2010):
	lon = math.Atan2(nvr.Y, -nvr.Z)

	// Equation (6) in Gade (2010) (Robust numerical solution)
	// vector component in the equatorial plane
	ec := math.Hypot(nvr.Y, nvr.Z)
	// atan() could also be used since latitude is within [-pi/2,pi/2]
	lat = math.Atan2(nvr.X, ec)

	// latitude = asin(n_E(1)) is a theoretical solution, but close to the Poles
	// it is ill-conditioned which may lead to numerical inaccuracies (and it will
	// give imaginary results for norm(n_E)>1)

	return lat, lon
}
