package options

import "gonum.org/v1/gonum/mat"

// Options contains optional parameters for n-vector calculations.
type Options struct {
	// CoordFrame defines the axes of the coordinate frame E.
	CoordFrame mat.Matrix
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
	}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
