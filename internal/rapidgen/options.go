package rapidgen

import (
	"github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/options"
	"pgregory.net/rapid"
)

// Options creates a rapid generator for n-vector options.
func Options() *rapid.Generator[[]options.Option] {
	return rapid.Custom(func(t *rapid.T) []options.Option {
		opts := make([]options.Option, 0)

		if rapid.Bool().Draw(t, "hasCoordFrame") {
			opts = append(
				opts,
				nvector.WithCoordFrame(RotationMat().Draw(t, "coordFrame")),
			)
		}

		if rapid.Bool().Draw(t, "hasEllipsoid") {
			opts = append(
				opts,
				nvector.WithEllipsoid(Ellipsoid().Draw(t, "ellipsoid")),
			)
		}

		return opts
	})
}
