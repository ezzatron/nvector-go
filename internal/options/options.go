package options

import (
	"gonum.org/v1/gonum/mat"
)

// Options contains optional parameters for n-vector calculations.
type Options struct {
	// CoordFrame defines the axes of the coordinate frame.
	CoordFrame mat.Matrix

	// Ellipsoid is the ellipsoid to use for calculations.
	Ellipsoid Ellipsoid
}

// Option is a functional option for n-vector calculations.
type Option func(*Options)

// New creates a new Options with the given options.
func New(opts []Option) *Options {
	o := &Options{
		CoordFrame: mat.NewDense(3, 3, []float64{
			0, 0, 1,
			0, 1, 0,
			-1, 0, 0,
		}),
		Ellipsoid: Ellipsoid{
			SemiMajorAxis: 6378137,
			Flattening:    1 / 298.257223563,
		},
	}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

// Ellipsoid is a reference ellipsoid.
type Ellipsoid struct {
	// SemiMajorAxis is the semi-major axis of the ellipsoid.
	SemiMajorAxis float64

	// Flattening is the flattening of the ellipsoid.
	Flattening float64
}
