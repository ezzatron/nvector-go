package rapidgen

import (
	"github.com/ezzatron/nvector-go"
	"github.com/ezzatron/nvector-go/internal/options"
	"pgregory.net/rapid"
)

// Options creates a rapid generator for n-vector options.
func Options() *rapid.Generator[[]options.Option] {
	return rapid.Custom(func(t *rapid.T) []options.Option {
		return rapid.SliceOfN(
			rapid.Just(
				nvector.WithCoordFrame(RotationMatrix().Draw(t, "coordFrame")),
			),
			0,
			1,
		).Draw(t, "opts")
	})
}
