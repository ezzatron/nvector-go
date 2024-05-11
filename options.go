package nvector

import (
	"gonum.org/v1/gonum/mat"

	"github.com/ezzatron/nvector-go/ellipsoid"
	"github.com/ezzatron/nvector-go/internal/options"
)

// Option is a functional option for n-vector calculations.
type Option = options.Option

// WithEllipsoid sets the ellipsoid option.
func WithEllipsoid(Ellipsoid ellipsoid.Ellipsoid) Option {
	return func(o *options.Options) {
		o.Ellipsoid = Ellipsoid
	}
}

// WithCoordFrame sets the coordinate frame option.
func WithCoordFrame(CoordFrame mat.Matrix) Option {
	return func(o *options.Options) {
		o.CoordFrame = CoordFrame
	}
}
