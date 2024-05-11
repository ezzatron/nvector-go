package nvector

import (
	"math"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/mat"
)

// FromLatLon converts a geodetic latitude and longitude to an n-vector.
//
// lat and lon are given in radians. The ellipsoid has no effect on the
// calculation.
//
// See: https://github.com/FFI-no/n-vector/blob/82d749a67cc9f332f48c51aa969cdc277b4199f2/nvector/lat_long2n_E.m
func FromLatLon(lat, lon float64, opts ...Option) mat.Vector {
	o := options.New(opts)

	// Equation (3) from Gade (2010):
	cosLat := math.Cos(lat)

	// CoordFrame selects correct E-axes
	nv := &mat.VecDense{}
	nv.MulVec(o.CoordFrame.T(), mat.NewVecDense(3, []float64{
		math.Sin(lat),
		math.Sin(lon) * cosLat,
		-math.Cos(lon) * cosLat,
	}))

	return nv
}

// ToLatLon converts an n-vector to a geodetic latitude and longitude.
//
// lat and lon are returned in radians. The ellipsoid has no effect on the
// calculation.
//
// See: https://github.com/FFI-no/n-vector/blob/82d749a67cc9f332f48c51aa969cdc277b4199f2/nvector/n_E2lat_long.m
func ToLatLon(nv mat.Vector, opts ...Option) (lat float64, lon float64) {
	o := options.New(opts)

	nvr := &mat.VecDense{}
	nvr.MulVec(o.CoordFrame, nv)

	// Equation (5) in Gade (2010):
	lon = math.Atan2(nvr.AtVec(1), -nvr.AtVec(2))

	// Equation (6) in Gade (2010) (Robust numerical solution)
	// vector component in the equatorial plane
	ec := math.Hypot(nvr.AtVec(1), nvr.AtVec(2))
	// atan() could also be used since latitude is within [-pi/2,pi/2]
	lat = math.Atan2(nvr.AtVec(0), ec)

	// latitude = asin(n_E(1)) is a theoretical solution, but close to the Poles
	// it is ill-conditioned which may lead to numerical inaccuracies (and it will
	// give imaginary results for norm(n_E)>1)

	return lat, lon
}
