package nvector

import (
	"gonum.org/v1/gonum/mat"

	"github.com/ezzatron/nvector-go/internal/options"
)

// Option is a functional option for n-vector calculations.
type Option = options.Option

// WithCoordFrame sets the coordinate frame option.
func WithCoordFrame(CoordFrame mat.Matrix) Option {
	return func(o *options.Options) {
		o.CoordFrame = CoordFrame
	}
}
