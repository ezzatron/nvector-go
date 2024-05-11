package nvector

import (
	"math"

	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/mat"
)

// FromLatLon converts a geodetic latitude and longitude to an n-vector.
//
// lat and lon are given in radians.
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
