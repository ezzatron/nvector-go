package nvector

import (
	"github.com/ezzatron/nvector-go/ellipsoid"
	"github.com/ezzatron/nvector-go/internal/options"
	"gonum.org/v1/gonum/spatial/r3"
)

// Option is a functional option for n-vector calculations.
type Option = options.Option

// WithEllipsoid sets the ellipsoid option.
func WithEllipsoid(e ellipsoid.Ellipsoid) Option {
	return func(o *options.Options) {
		o.Ellipsoid = e
	}
}

// WithCoordFrame sets the coordinate frame option.
func WithCoordFrame(f *r3.Mat) Option {
	return func(o *options.Options) {
		o.CoordFrame = f
	}
}
